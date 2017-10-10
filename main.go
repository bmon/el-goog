package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()

	// handlers for index.html and bundle.js
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/dist/index.html")
	})
	r.HandleFunc("/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/dist/bundle.js")
	})
	r.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/dist/TEMPLOGO.png")
	})

	// Bind to a port and pass our router in
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":5000", loggedRouter))
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
