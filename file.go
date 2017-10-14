package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type File struct {
	ID       int
	Parent   *Folder
	Name     string
	Size     int
	Checksum string
	Modified time.Time
}

func CreateFile(name string, size int, checksum string, parent *Folder) *File {
	f := &File{0, parent, name, size, checksum, time.Now()}
	f.Insert()
	return f
}

func (f *File) Insert() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "insert into files values (NULL, $1, $2, $3, $4, $5)"
	res, err := db.Exec(sqlStmt, f.Parent, f.Name, f.Size, f.Checksum, f.Modified.Unix())
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

	sqlStmt := "update folders set parent_id=$1, name=$2, size=$3, checksum=$4, modified=$5 where id=$6"
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

	/*
	 * Dont make the same mistake i did:
	 * Documentation for fine uploader is on their website
	 * but not in Docs, in API.
	 *
	 */
	user := GetRequestUser(r)
	if user == nil {
		fmt.Println("user not authenticated!")
		http.Error(w, "You must be authenticated to perform this action", 401)
		return
	} else {
		fmt.Println("user authenticated!", user.Email)
	}

	file, header, err := r.FormFile("qqfile")

	// TODO appropriately utilise these
	fmt.Println(file)   // multipart.File, which happens to point to an interface?
	fmt.Println(header) // also has the body
	fmt.Println(err)

	if err == nil {
		fmt.Fprintf(w, "{\"success\":true}")
	} else {
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
}
