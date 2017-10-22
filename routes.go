package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Resource    string
	HandlerFunc http.HandlerFunc
}

type Op struct {
	Method  string
	Handler http.HandlerFunc
}

type Resource struct {
	URI string
	Ops []Op
}

type Routes []Resource

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, resource := range routes {
		for _, op := range resource.Ops {
			router.
				Path(resource.URI).
				Methods(op.Method).
				Handler(op.Handler)
		}
	}
	return router
}

var routes = Routes{
	Resource{
		"/users",
		[]Op{{"POST", UserCreateHandler}},
	},
	Resource{
		"/users/{id:[0-9]+}",
		[]Op{
			{"GET", UserGetDetails},
			{"DELETE", UserDelete},
			{"PUT", UserModifyHandler},
		},
	},
	Resource{
		"/folders/",
		[]Op{{"POST", FolderCreateHandler}},
	},
	Resource{
		"/folders/{id:[0-9]+}",
		[]Op{
			{"GET", FolderGetHandler},
			{"DELETE", FolderDeleteHandler},
		},
	},
	Resource{
		"/files",
		[]Op{
			{"GET", FileGetHandler},
		},
	},
	Resource{
		"/folders/{id:[0-9]+}/files",
		[]Op{
			{"POST", FileCreateHandler},
		},
	},
	Resource{
		"/files/{id:[0-9]+}",
		[]Op{
			{"GET", FileGetHandler},
			{"DELETE", FileDeleteHandler},
		},
	},
	Resource{
		"/login",
		[]Op{{"POST", UserLogin}},
	},
	Resource{
		"/logout",
		[]Op{{"GET", UserLogout}},
	},
}
