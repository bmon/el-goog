package main

import (
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123")
}
