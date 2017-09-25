package main

import (
	"fmt"
	"net/http"
)

type User struct {
	id       int
	email    string
	password string
	verified bool
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123")
}
