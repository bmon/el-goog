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
		db, err := sql.Open("sqlite3", "./elgoog.db")
	        if  err != nil {
                        fmt.Fprintf(w, "ERROR: could not access database")
			fmt.Println("couldn't open le database")
		}
		stmt, err := db.Prepare("INSERT INTO users(id, email, password, username) values(?,?,?,?)")
	        if  err != nil {
                        fmt.Fprintf(w, "ERROR: faulty SQL command")
                        fmt.Println("insert line is broken, I dun goof'd") 
                }
		res, err := stmt.Exec(nil, eml, pwd, usr)
	        if  err != nil {
                        fmt.Fprintf(w, "ERROR: Failed to execute SQL insert")
                        fmt.Println("could not insert lel") 
                }
		id, err := res.LastInsertId()
		if err !=nil {
                        fmt.Fprintf(w, "ERROR: failed to add data to database")
			fmt.Println("There was an insert errror")
		} else {
			fmt.Println("Result of SQL insert: "+strconv.FormatInt(id, 10))
		}
                fmt.Fprintf(w, "Success! ID is:"+strconv.FormatInt(id,10))
	} else {
		fmt.Fprintf(w, "ERROR: invalid email address")
	        fmt.Println("yeah nah mate get a better email address")
	}
}
