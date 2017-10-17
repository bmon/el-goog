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

	sqlStmt := "update folders set parent_id=?, name=?, size=?, modified=? where id=?"
	_, err = db.Exec(sqlStmt, f.Parent, f.Name, f.Modified.Unix(), f.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
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

	// =======================================================================
	// Folder Verification
	vars := mux.Vars(r)
	folderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "Malformed folder ID"}`, 400)
		fmt.Println("bad a to i with folder id:" + vars["id"])
		return
	}
	rootFolder, err := FolderSelectByID(folderID)
	if err != nil {
		http.Error(w, `{"error": "Invalid folder ID"}`, 400)
		fmt.Println("folder select by id failed with id:" + vars["id"])
		return
	}
	rootFolder.GetUserID()
	if user.ID != rootFolder.GetUserID() {
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

		// =================================================================
		//  Folders
		// was partly validated before the nested if's

		filePath := "/tmp" + rootFolder.Path()
		fmt.Println("downloading file to filePath == " + filePath)
		os.MkdirAll(filePath, 0755)

		// =================================================================
		//  File
		//var osFile os.File

		if part == 0 {
			db, err := sql.Open("sqlite3", DatabaseFile)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			// make a database entry
			dbFile := CreateFile(fileName,
				0,          /*size*/
				rootFolder, /*TODO client path*/
			)
			dbFile.Insert()
		}

		// create does return the osFile,
		// however golang much prefers the types are strict here
		//osFile, err := os.Create(rootFolderDir + `/` + fileName)
		osFile, err := os.OpenFile(filePath+`/`+fileName, os.O_RDWR|os.O_CREATE, 0755)

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
					if part == totalParts-1 {

						db, err := sql.Open("sqlite3", DatabaseFile)
						if err != nil {
							log.Fatal(err)
						}
						defer db.Close()

						sqlStmt := "update files set size = ? where id = ?"
						_, err = db.Exec(sqlStmt, totalSize, 0)
						if err != nil {
							retry = 1
							errMsg += `"db": "failed"`
							fmt.Println(errMsg)
						}
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

func (f *File) Path() string {
	return f.Parent.Path() + fmt.Sprintf("%s.%d", f.Name, f.ID)
}
