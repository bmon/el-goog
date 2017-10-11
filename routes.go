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
		"UserCreate",
		"POST",
		"/user/create",
		UserCreate,
	},
	Route{
		"UserLogin",
		"POST",
		"/user/login",
		UserLogin,
	},
	Route{
		"UserLogout",
		"POST",
		"/user/logout",
		UserLogout,
	},
	Route{
		"UploadHandler",
		"POST",
		"/upload",
		UploadHandler,
	},
}
