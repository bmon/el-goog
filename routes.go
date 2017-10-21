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
		"/folders/{id}/files",
		FileCreateHandler,
	},
	Route{
		"FileGetHandler",
		"GET",
		"/files/{id:[0-9]+}",
		FileGetHandler,
	},
	Route{
		"FolderGetHandler",
		"GET",
		"/folders/{id:[0-9]+}",
		FolderGetHandler,
	},
	Route{
		"usersDelete",
		"DELETE",
		"/users/{id:[0-9]+}",
		UserDelete,
	},
	Route{
		"FolderDeleteHandler",
		"DELETE",
		"/folders/{id:[0-9]+}",
		FolderDeleteHandler,
	},
	Route{
		"FileDeleteHandler",
		"DELETE",
		"/files/{id:[0-9]+}",
		FileDeleteHandler,
	},
        Route{
                "UserGetDetails",
                "GET",
                "/users/{id:[0-9]+}",
                UserGetDetails,
        },
        Route{
                "UserModifyHandler",
                "PUT",
                "/users/{id:[0-9]+}",
                UserModifyHandler,
        },
}
