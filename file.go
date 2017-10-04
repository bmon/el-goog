package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
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
