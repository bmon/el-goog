package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
	"unicode"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Folder struct {
	ID       int       `json:"id"`
	Parent   *Folder   `json:"-"`
	Name     string    `json:"name"`
	Modified time.Time `json:"modified"`
}

type SerialFolder struct {
	ID           int       `json:"id"`
	ParentID     int       `json:"parent_id"`
	Name         string    `json:"name"`
	Modified     time.Time `json:"modified"`
	Path         string    `json:"path"`
	ChildFolders []Folder  `json:"child_folders"`
	ChildFiles   []File    `json:"child_files"`
}

func CreateFolder(name string, parent *Folder) *Folder {
	f := &Folder{0, parent, name, time.Now()}
	f.Insert()
	return f
}

func (f *Folder) MakeSerial(query string) SerialFolder {
	if query == "" {
		query = "%"
	} else {
		query = fmt.Sprintf("%%%s%%", query)
	}
	parentID := -1
	if f.Parent != nil {
		parentID = f.Parent.ID
	}
	folder := SerialFolder{f.ID, parentID, f.Name, f.Modified, f.Path(), make([]Folder, 0), make([]File, 0)}

	rows, err := DB.Query("SELECT id FROM folders WHERE parent_id=? and name like ?", folder.ID, query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			child, _ := FolderSelectByID(id)
			folder.ChildFolders = append(folder.ChildFolders, *child)
		} else {
			fmt.Println(err)
		}
	}

	rows, err = DB.Query("SELECT id FROM files WHERE parent_id=? and name like ?", folder.ID, query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			child, _ := FileSelectByID(id)
			folder.ChildFiles = append(folder.ChildFiles, *child)
		} else {
			fmt.Println(err)
		}
	}
	return folder
}

func FolderSelectByID(folderID int) (*Folder, error) {
	f := &Folder{}
	var timestamp int64
	err := DB.QueryRow("SELECT id, name, modified, parent_id FROM folders WHERE id=?", folderID).Scan(&f.ID, &f.Name, &timestamp, &f.Parent)
	if err != nil {
		return nil, err
	}
	f.Modified = time.Unix(timestamp, 0)
	return f, nil
}

func (f *Folder) Insert() {
	sqlStmt := "insert into folders values (NULL, ?, ?, ?)"
	res, err := DB.Exec(sqlStmt, f.Parent, f.Name, f.Modified.Unix())
	if err != nil {
		fmt.Println("insert error", err)
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
		} else {
			f.ID = int(id)
		}
	}
}

func (f *Folder) Update() {
	f.Modified = time.Now()
	sqlStmt := "update folders set parent_id=?, name=?, modified=? where id=?"
	_, err := DB.Exec(sqlStmt, f.Parent, f.Name, f.Modified.Unix(), f.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
}

func (f *Folder) GetUserID() int {
	root := f
	for root.Parent != nil {
		root = root.Parent
	}

	var userID int
	err := DB.QueryRow("select id from users where root_folder = ?", root.ID).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return userID
}

// This method allows us to do db.exec with a folder instance argument
func (f *Folder) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return int64(f.ID), nil
}

func (f *Folder) Scan(value interface{}) error {
	if value == nil {
		fmt.Println("value is nil!!")
	}
	if id, err := driver.Int32.ConvertValue(value); err == nil {
		if v, ok := id.(int64); ok {
			newf, err := FolderSelectByID(int(v))
			if err == nil {
				f.ID = newf.ID
				f.Name = newf.Name
				f.Parent = newf.Parent
				f.Modified = newf.Modified
			} else {
				f.Parent = nil
			}
		}
	} else {
		return err
	}
	return nil
}

func (f *Folder) Path() string {
	dirname := fmt.Sprintf("%s.%d/", f.Name, f.ID)
	var parent string
	if f.Parent != nil {
		parent = f.Parent.Path()
	} else {
		user, err := UserSelectByID(f.GetUserID())
		if err == nil {
			parent = fmt.Sprintf("uploads/%s/", user.Email)
		}
	}
	return parent + dirname
}

func FolderGetHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
	vars := mux.Vars(r)
	vals := r.URL.Query()

	folderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	f, err := FolderSelectByID(folderID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if user == nil || user.ID != f.GetUserID() {
		http.Error(w, "You do not have permission to retrieve this object", 403)
	}

	sortVal := vals.Get("sort")
	s := f.MakeSerial(vals.Get("q"))
	// pick the appropriate sort function
	switch sortVal {

	case "-name":
		sort.Slice(s.ChildFiles, func(i, j int) bool {
			return unicode.ToLower(rune(s.ChildFiles[i].Name[0])) > unicode.ToLower(rune(s.ChildFiles[j].Name[0]))
		})
		sort.Slice(s.ChildFolders, func(i, j int) bool {
			return unicode.ToLower(rune(s.ChildFolders[i].Name[0])) > unicode.ToLower(rune(s.ChildFolders[j].Name[0]))
		})
	case "size":
		sort.Slice(s.ChildFiles, func(i, j int) bool { return s.ChildFiles[i].Size < s.ChildFiles[j].Size })
	case "-size":
		sort.Slice(s.ChildFiles, func(i, j int) bool { return s.ChildFiles[i].Size > s.ChildFiles[j].Size })
	case "modified":
		sort.Slice(s.ChildFiles, func(i, j int) bool { return s.ChildFiles[i].Modified.Unix() < s.ChildFiles[j].Modified.Unix() })
		sort.Slice(s.ChildFolders, func(i, j int) bool { return s.ChildFolders[i].Modified.Unix() < s.ChildFolders[j].Modified.Unix() })
	case "-modified":
		sort.Slice(s.ChildFiles, func(i, j int) bool { return s.ChildFiles[i].Modified.Unix() > s.ChildFiles[j].Modified.Unix() })
		sort.Slice(s.ChildFolders, func(i, j int) bool { return s.ChildFolders[i].Modified.Unix() > s.ChildFolders[j].Modified.Unix() })
	default: // name
		sort.Slice(s.ChildFiles, func(i, j int) bool {
			return unicode.ToLower(rune(s.ChildFiles[i].Name[0])) < unicode.ToLower(rune(s.ChildFiles[j].Name[0]))
		})
		sort.Slice(s.ChildFolders, func(i, j int) bool {
			return unicode.ToLower(rune(s.ChildFolders[i].Name[0])) < unicode.ToLower(rune(s.ChildFolders[j].Name[0]))
		})
	}

	res, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(res)
}

func FolderCreateHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
	r.ParseForm()
	name := r.Form.Get("name")
	parent := r.Form.Get("parent")

	folderID, err := strconv.Atoi(parent)
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	f, err := FolderSelectByID(folderID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if user == nil || user.ID != f.GetUserID() {
		http.Error(w, "You do not have permission to retrieve this object", 403)
		return
	}
	new := CreateFolder(name, f)
	res, err := json.MarshalIndent(new, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(res)

	return
}

func FolderDeleteHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
	vars := mux.Vars(r)
	folderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	f, err := FolderSelectByID(folderID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if user.ID != f.GetUserID() {
		http.Error(w, "You do not have permission to retrieve this object", 403)
		return
	}
	f.Delete()
}

func (f *Folder) Delete() {
	serial := f.MakeSerial("")
	for _, file := range serial.ChildFiles {
		file.Delete()
	}
	for _, folder := range serial.ChildFolders {
		folder.Delete()
	}
	DB.Exec("DELETE from folders where id=?", f.ID)
}
