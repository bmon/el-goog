package main

import (
	"fmt"
	"net/http"
	//"database/sql"
	"regexp"
)

type User struct {
	id       int
	email    string
	password string
	verified bool
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	fmt.Println("Here's the password so go stops complaining: "+password)
	
	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)

	if(isEmail.MatchString(email)) {
		fmt.Fprintf(w, "yeah mate is valid")
        	fmt.Println("yeah mate is valid")
	} else {
		fmt.Fprintf(w, "yeah nah mate get a better email address")
	        fmt.Println("yeah nah mate get a better email address")
	}
}
