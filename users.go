package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var pwdsalt []byte = []byte(getEnv("PASSWORD_SALT", "ayy-lmao_top-kek_meme"))

type User struct {
	ID         int
	Email      string
	Password   string
	Username   string
	RootFolder *Folder
}

func (u *User) Insert() error {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := "INSERT INTO users(id, email, password, username, root_folder) values(?,?,?,?,?)"
	res, err := db.Exec(sqlStmt, nil, u.Email, u.Password, u.Username, u.RootFolder.ID)
	if err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			u.ID = int(id)
		}
	}
	return nil
}

func UserSelectByID(userID int) (*User, error) {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return nil, err
	}

	u := &User{}
	rootID := -1
	err = db.QueryRow("SELECT id, email, password, username, root_folder FROM users WHERE id=?", userID).Scan(&u.ID, &u.Email, &u.Password, &u.Username, &rootID)
	if err != nil {
		return nil, err
	}
	u.RootFolder, err = FolderSelectByID(rootID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	username := r.PostFormValue("username")

	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)

	if isEmail.MatchString(email) {

		sltpwd := append([]byte(password), pwdsalt...)
		hshpwd, _ := bcrypt.GenerateFromPassword(sltpwd, 10) //salting and hashing the password

		hashedPassword := string(hshpwd[:])

		user := &User{-1, email, hashedPassword, username, CreateFolder("root", nil)}
		err := user.Insert()

		if err != nil {
			user.RootFolder.Delete()
			if err, ok := err.(sqlite3.Error); ok {
				if err.Code == sqlite3.ErrConstraint {
					http.Error(w, "Email already in use.", 400)
					return
				}
			}

			fmt.Println("ERROR:", err)
			http.Error(w, err.Error(), 500)
			return
		}
		user.CreateSession(w)

		fmt.Fprintf(w, "Success! ID is: %d", user.ID)
	} else {
		http.Error(w, "ERROR: invalid email address. Received address: "+email, 400)
		return
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if GetRequestUser(r) != nil {
		http.Error(w, "Already logged in", 400)
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// open
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// run it
	var dbPass string
	var userID int
	err = db.QueryRow("SELECT id,password FROM users WHERE email = ?", email).Scan(&userID, &dbPass)
	if err != nil {
		http.Error(w, "Wrong username or password", 400)
		return
	}

	sltpwd := append([]byte(password), pwdsalt...)
	err = bcrypt.CompareHashAndPassword([]byte(dbPass), sltpwd)

	if err != nil {
		http.Error(w, "Wrong username or password", 400)
		return
	}
	user, err := UserSelectByID(userID)
	if err != nil {
		//this should never happen -- we just verified the user exists
		http.Error(w, err.Error(), 500)
	}
	user.CreateSession(w)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	DeleteRequestSession(w, r)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	thisUser := GetRequestUser(r)

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	if thisUser.ID != userID {
		http.Error(w, "You do not have permission to modify this account", 403)
		return
	}

	thisUser.RootFolder.Delete()

	DeleteRequestSession(w, r)

	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return
	}
	defer db.Close()

	sqlStmt := "DELETE FROM users WHERE id=?"
	_, err = db.Exec(sqlStmt, thisUser.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
}

func UserGetDetails(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)

	vars := mux.Vars(r)
        userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println(err)
                http.NotFound(w, r)
                return
	}

	if user.ID != userID {
                http.Error(w, "You do not have permission to view this information", 403)
                return
        }

        fmt.Fprintf(w, "username:%s\n", user.Username)
        fmt.Fprintf(w, "email:%s", user.Email)
}

func UserModifyHandler(w http.ResponseWriter, r *http.Request) {

}
