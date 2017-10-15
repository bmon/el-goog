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

	sqlStmt := "insert into folders values (?, ?, ?, ?)"
	res, err := db.Exec(sqlStmt, f.ID, f.Parent, f.Name, f.Modified.Unix())
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
		parent = fmt.Sprintf("uploads/userEmailTODO/")
	}
	return parent + dirname
}
