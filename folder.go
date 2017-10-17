package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	ChildFolders []Folder  `json:"child_folders"`
	ChildFiles   []File    `json:"child_files"`
}

func CreateFolder(name string, parent *Folder) *Folder {
	f := &Folder{0, parent, name, time.Now()}
	f.Insert()
	return f
}

func (f *Folder) MakeSerial() SerialFolder {
	folder := SerialFolder{f.ID, f.Parent.ID, f.Name, f.Modified, make([]Folder, 0), make([]File, 0)}
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT id FROM folders WHERE parent_id=?", folder.ID)
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

	rows, err = db.Query("SELECT id FROM files WHERE parent_id=?", folder.ID)
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
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return nil, err
	}

	f := &Folder{}
	var timestamp int64
	err = db.QueryRow("SELECT id, name, modified, parent_id FROM folders WHERE id=?", folderID).Scan(&f.ID, &f.Name, &timestamp, &f.Parent)
	if err != nil {
		return nil, err
	}
	f.Modified = time.Unix(timestamp, 0)
	return f, nil
}

func (f *Folder) Insert() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "insert into folders values (NULL, ?, ?, ?)"
	res, err := db.Exec(sqlStmt, f.Parent, f.Name, f.Modified.Unix())
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

func (f *Folder) Delete() {
	f.Modified = time.Now()
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "delete from folders where id=?"
	_, err = db.Exec(sqlStmt, f.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func (f *Folder) Update() {
	f.Modified = time.Now()
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "update folders set parent_id=?, name=?, modified=? where id=?"
	_, err = db.Exec(sqlStmt, f.Parent, f.Name, f.Modified.Unix(), f.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
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
	if f.ID == 0 {
		fmt.Printf("ERROAR %+v\n", f)
	}
	dirname := fmt.Sprintf("%s.%d/", f.Name, f.ID)
	var parent string
	if f.Parent != nil {
		parent = f.Parent.Path()
	} else {
		db, err := sql.Open("sqlite3", DatabaseFile)
		if err != nil {
			return "ERROR"
		}
		defer db.Close()
		row := db.QueryRow("SELECT email FROM users WHERE root_folder=?", f.ID)
		email := ""
		err = row.Scan(&email)
		parent = fmt.Sprintf("uploads/%s/", email)
	}
	return parent + dirname
}

//Test function for Path()
func FolderPath(w http.ResponseWriter, r *http.Request) {
	fID := r.PostFormValue("id")
	folderID, _ := strconv.Atoi(fID)
	f, err := FolderSelectByID(folderID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	path := f.Path()
	fmt.Fprintf(w, path)
}

func FolderGetHandler(w http.ResponseWriter, r *http.Request) {
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
	res, err := json.MarshalIndent(f.MakeSerial(), "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(res)
}
