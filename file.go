package main

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type File struct {
	ID       int       `json:"id"`
	Parent   *Folder   `json:"-"`
	Name     string    `json:"name"`
	Size     int       `json:"size"`
	Modified time.Time `json:"modified"`
}

func CreateFile(name string, size int, parent *Folder) *File {
	f := &File{0, parent, name, size, time.Now()}
	f.Insert()
	return f
}

func FileSelectByID(fileID int) (*File, error) {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return nil, err
	}

	f := &File{}
	var timestamp int64
	err = db.QueryRow("SELECT id, name, size, modified, parent_id FROM files WHERE id=?", fileID).Scan(&f.ID, &f.Name, &f.Size, &timestamp, &f.Parent)
	if err != nil {
		return nil, err
	}
	f.Modified = time.Unix(timestamp, 0)
	return f, nil
}

func (f *File) Insert() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "insert into files values (NULL, ?, ?, ?, ?)"
	res, err := db.Exec(sqlStmt, f.Parent, f.Name, f.Size, f.Modified.Unix())
	if err != nil {
		fmt.Println("file insert error", err)
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
		} else {
			f.ID = int(id)
		}
	}
}

func (f *File) Update() {
	f.Modified = time.Now()
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "update files set parent_id=?, name=?, size=?, modified=? where id=?"
	_, err = db.Exec(sqlStmt, f.Parent.ID, f.Name, f.Size, f.Modified.Unix(), f.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
}

func (f *File) GetUserID() int {
	return f.Parent.GetUserID()
}

// This method allows us to do db.exec with a folder instance argument
func (f *File) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return int64(f.ID), nil
}

func FileCreateHandler(w http.ResponseWriter, r *http.Request) {

	retry := 0
	errMsg := ""

	user := GetRequestUser(r)
	if user == nil {
		http.Error(w, `{"error": "You must be authenticated to perform this action"}`, 401)
		return
	}

	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// =======================================================================
	// Folder Verification
	vars := mux.Vars(r)
	folderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Malformed folder ID"}`, 400)
		fmt.Println("bad a to i with folder id:" + vars["id"])
		return
	}
	targetFolder, err := FolderSelectByID(folderID)
	if err != nil {
		http.Error(w, `{"error": "Invalid folder ID"}`, 400)
		fmt.Println("folder select by id failed with id:" + vars["id"])
		return
	}
	targetFolder.GetUserID()
	if user.ID != targetFolder.GetUserID() {
		http.Error(w, `{"error": "You must be the owner of this directory to perform this action"}`, 401)
		return
	}
	// =======================================================================
	// formfile
	httpFile, _ /*header*/, err := r.FormFile("qqfile")
	defer httpFile.Close()

	if err != nil {
		retry = 1
		errMsg += `,"qqfile": "failed"`
		fmt.Println(errMsg)
	} else {
		// extract
		part, _ := strconv.ParseInt(r.PostFormValue("qqpartindex"), 10, 64)
		offset, _ := strconv.ParseInt(r.PostFormValue("qqpartbyteoffset"), 10, 64)
		//size, _ := strconv.ParseInt(r.PostFormValue("qqchunksize"), 10, 64)
		totalSize, _ := strconv.ParseInt(r.PostFormValue("qqtotalfilesize"), 10, 64)
		totalParts, _ := strconv.ParseInt(r.PostFormValue("qqtotalparts"), 10, 64)
		fileName := r.PostFormValue("qqfilename")

		//  File
		//var osFile os.File

		var dbFile *File
		if part == 0 {
			// make a database entry
			dbFile = CreateFile(fileName,
				-1,           /*size*/
				targetFolder, /*TODO client path*/
			)
		} else {
			var fileID int
			err = db.QueryRow("select id from files where size=-1 and parent_id=? and name=?", folderID, fileName).Scan(&fileID)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			dbFile, err = FileSelectByID(fileID)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}

		// =================================================================
		//  Folders
		// was partly validated before the nested if's

		filePath := "./tmp/" + dbFile.Path()
		fmt.Println("downloading file to filePath == " + filePath)
		os.MkdirAll("./tmp/"+dbFile.Parent.Path(), 0755)

		// =================================================================

		// create does return the osFile,
		// however golang much prefers the types are strict here
		//osFile, err := os.Create(rootFolderDir + `/` + fileName)
		osFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

		// =================================================================
		//  Writing
		if err != nil {
			retry = 1
			errMsg += `,"osfile": "failed"`
			fmt.Println(errMsg)
		} else {
			// copy the http file to the os file
			defer osFile.Close()
			b := bytes.NewBuffer(nil)
			_ /*bRead*/, err := io.Copy(b, httpFile)

			if err != nil {
				retry = 1
				errMsg += `,"iocopy": "failed"`
				fmt.Println(errMsg)
			} else {
				// write out to the file
				_ /*bWritten*/, err := osFile.WriteAt(b.Bytes(), offset)

				if err != nil {
					retry = 1
					errMsg += `,"oswrite": "failed"`
					fmt.Println(errMsg)
					fmt.Println(err)
				} else {
					// final chunk: also update database (size, its finished) etc.
					fmt.Println(part, totalParts, "asdgjksadgafadsjgd")
					if part == totalParts-1 {
						dbFile.Size = int(totalSize)
						dbFile.Update()
					}
				}
			}
		}
	}
	// =======================================================================

	// handle success/ notify accumulated errors (so client may tell user someday)
	if retry == 0 {
		fmt.Fprintf(w, `{"success":true}`)
	} else {
		fmt.Fprintf(w, `{"error": "%v"`+errMsg+`}`, err)
	}
}

func FileGetHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
        if user == nil {
                http.Error(w, "User is not logged in", 403)
		return
        }
	vars := mux.Vars(r)
	fileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	f, err := FileSelectByID(fileID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if user.ID != f.Parent.GetUserID() {
		http.Error(w, "You do not have permission to retrieve this object", 403)
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+f.Name)
	http.ServeFile(w, r, "./tmp/"+f.Path())
}

func (f *File) Path() string {
	return f.Parent.Path() + fmt.Sprintf("%s.%d", f.Name, f.ID)
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
        if user == nil {
                http.Error(w, "User is not logged in", 403)
		return
        }
	vars := mux.Vars(r)
	fileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	f, err := FileSelectByID(fileID)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if user.ID != f.Parent.GetUserID() {
		http.Error(w, "You do not have permission to retrieve this object", 403)
		return
	}
	f.Delete()
}

func (f *File) Delete() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("DELETE from files where id=?", f.ID)
}

func FilesGetHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)
        if user == nil {
                http.Error(w, "User is not logged in", 403)
		return
        }

	query := r.URL.Query()
	search := query.Get("q")

	db, err := sql.Open("sqlite3", DatabaseFile)
        if err != nil {
                fmt.Println(err)
        }
	defer db.Close()

	root := user.RootFolder
	sFolder := SerialFolder{root.ID, -1, root.Name, root.Modified, root.Path(), make([]Folder, 0), make([]File, 0)}

	toSearch := "%"+search+"%"

	rows, err := db.Query("SELECT * FROM files WHERE name LIKE ?", toSearch)
        if err != nil {
                fmt.Println(err)
        }
        defer rows.Close()

        for rows.Next() {
                f := &File{}
		var pid int
		var timeStamp int64
                if err := rows.Scan(&f.ID, &pid, &f.Name, &f.Size, &timeStamp); err == nil {
			f.Parent, err = FolderSelectByID(pid)
			f.Modified = time.Unix(timeStamp, 0)
			if err != nil{
				http.NotFound(w, r)
				fmt.Println(err)
		                return
			}
                        if f.GetUserID() == user.ID {
				sFolder.ChildFiles = append(sFolder.ChildFiles, *f)
			}
                } else {
                        fmt.Println(err)
			return
                }
        }
	res, err := json.MarshalIndent(sFolder, "", "\t")
        if err != nil {
                http.Error(w, err.Error(), 500)
        }
        w.Write(res)
}
