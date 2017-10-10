package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
	"strconv"
)

type User struct {
	id       int
	email    string
	password string
	verified bool
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	eml := r.PostFormValue("email")
	pwd := r.PostFormValue("password")
        usr := r.PostFormValue("username")
	
	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)

	if(isEmail.MatchString(eml)) {
		fmt.Fprintf(w, "yeah mate is valid email")
        	
		db, err := sql.Open("sqlite3", "./elgoog.db")
	        if  err != nil {
			fmt.Println("couldn't open le database")
		}
		stmt, err := db.Prepare("INSERT INTO users(id, email, password, username) values(?,?,?,?)")
	        if  err != nil {
                        fmt.Println("insert line is broken, I dun goof'd") 
                }
		res, err := stmt.Exec(nil, eml, pwd, usr)
	        if  err != nil {
                        fmt.Println("could not insert lel") 
                }
		id, err := res.LastInsertId()
		if err !=nil {
			fmt.Println("There was an insert errror")
		} else {
			fmt.Println("Result of SQL insert: "+strconv.FormatInt(id, 10))
		}
	} else {
		fmt.Fprintf(w, "yeah nah mate get a better email address")
	        fmt.Println("yeah nah mate get a better email address")
	}
}
