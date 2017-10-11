package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id       int
	email    string
	password string
	verified bool
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	eml := r.PostFormValue("email")
	fmt.Println(eml)
	pwd := r.PostFormValue("password")
	fmt.Println(pwd)
	usr := r.PostFormValue("username")
	fmt.Println(usr)

	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)

	if isEmail.MatchString(eml) {
		db, err := sql.Open("sqlite3", "./elgoog.db")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		stmt, err := db.Prepare("INSERT INTO users(id, email, password, username) values(?,?,?,?)")
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		res, err := stmt.Exec(nil, eml, pwd, usr)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		} else {
			fmt.Println("Result of SQL insert: " + strconv.FormatInt(id, 10))
		}
		fmt.Fprintf(w, "Success! ID is:"+strconv.FormatInt(id, 10))
	} else {
		http.Error(w, "ERROR: invalid email address. Received address: "+eml, 400)
		return
	}
}
