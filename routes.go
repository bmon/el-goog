package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"usersCreate",
		"POST",
		"/users",
		UserCreate,
	},
	Route{
		"usersLogin",
		"POST",
		"/login",
		UserLogin,
	},
	Route{
		"usersLogout",
		"GET",
		"/logout",
		UserLogout,
	},
	Route{
		"FileCreateHandler",
		"POST",
		"/folders/{id:[0-9]+}/files",
		FileCreateHandler,
	},
	Route{
		"FolderGetHandler",
		"GET",
		"/folders/{id:[0-9]+}",
		FolderGetHandler,
	},
	Route{
		"FolderPath",
		"POST",
		"/path",
		FolderPath,
	},
	Route{
		"FilePath",
		"POST",
		"/filepath",
		FilePath,
	},
}
