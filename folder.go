package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Folder struct {
	ID       int
	Parent   *Folder
	Name     string
	Modified time.Time
}

func CreateFolder(name string, parent *Folder) *Folder {
	f := &Folder{0, parent, name, time.Now()}
	f.Insert()
	return f
}

func (f *Folder) Insert() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "insert into folders values (NULL, $1, $2, $3)"
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

func (f *Folder) Update() {
	f.Modified = time.Now()
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "update folders set parent_id=$1, name=$2, modified=$3 where id=$4"
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

func (f *Folder) Path() string {
	if f.Parent == nil {
		//DO SELECT LINE TO GET EMAIL ADDRESS
		return "uploads/" + /*[INSERT EMAIL HERE] +*/ "/" + f.Name + "." + fmt.Sprintf("%d", f.ID)
	} else {
		return f.Parent.Path() + "/" + f.Name + "." + fmt.Sprintf("%d", f.ID)
	}
}
