package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"encoding/json"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var pwdsalt []byte = []byte(getEnv("PASSWORD_SALT", "ayy-lmao_top-kek_meme"))

type User struct {
	ID         int     `json:"id"`
	Email      string  `json:"email"`
	Password   string  `json:"-"`
	Username   string  `json:"username"`
	RootFolder *Folder `json:"root_folder"`
}

func CreateUser(name, email, password string) (*User, error) {
	sltpwd := append([]byte(password), pwdsalt...)
	hshpwd, _ := bcrypt.GenerateFromPassword(sltpwd, 10) //salting and hashing the password

	hashedPassword := string(hshpwd[:])

	user := &User{-1, email, hashedPassword, name, CreateFolder("root", nil)}
	err := user.Insert()

	if err != nil {
		user.RootFolder.Delete()
		return nil, err
	}
	return user, nil
}

func (u *User) Insert() error {
	sqlStmt := "INSERT INTO users(id, email, password, username, root_folder) values(?,?,?,?,?)"
	res, err := DB.Exec(sqlStmt, nil, u.Email, u.Password, u.Username, u.RootFolder.ID)
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
	u := &User{}
	rootID := -1
	err := DB.QueryRow("SELECT id, email, password, username, root_folder FROM users WHERE id=?", userID).Scan(&u.ID, &u.Email, &u.Password, &u.Username, &rootID)
	if err != nil {
		return nil, err
	}
	u.RootFolder, err = FolderSelectByID(rootID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	username := r.PostFormValue("username")

	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)
	if isEmail.MatchString(email) {
		user, err := CreateUser(username, email, password)
		if err != nil {
			if err, ok := err.(sqlite3.Error); ok {
				if err.Code == sqlite3.ErrConstraint {
					http.Error(w, "Email already in use.", 400)
					return
				}
			}
			http.Error(w, err.Error(), 500)
			return
		}
		user.CreateSession(w)
		fmt.Fprintf(w, "Success! ID is: %d", user.ID)
		return
	}
	http.Error(w, "ERROR: invalid email address. Received address: "+email, 400)
	return
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if GetRequestUser(r) != nil {
		http.Error(w, "Already logged in", 400)
		return
	}
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// run it
	var dbPass string
	var userID int
	err := DB.QueryRow("SELECT id,password FROM users WHERE email = ?", email).Scan(&userID, &dbPass)
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

	if thisUser == nil {
		http.Error(w, "User is not logged in", 403)
		return
	}

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
	sqlStmt := "DELETE FROM users WHERE id=?"
	_, err = DB.Exec(sqlStmt, thisUser.ID)
	if err != nil {
		fmt.Println(sqlStmt, err)
		fmt.Println(err)
	}
}

func UserGetDetails(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)

	if user == nil {
		http.Error(w, "User is not logged in", 403)
		return
	}

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

	res, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Write(res)
}

func UserModifyHandler(w http.ResponseWriter, r *http.Request) {
	user := GetRequestUser(r)

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if user == nil || user.ID != userID {
		http.Error(w, "You do not have permission to modify this account", 403)
		return
	}

	newPwd := r.FormValue("password")
	username := r.FormValue("username")

	var hashedNewPassword string

	if newPwd != "" {
		hshNewPwd, err := bcrypt.GenerateFromPassword(append([]byte(newPwd), pwdsalt...), 10)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		hashedNewPassword = string(hshNewPwd[:])
	} else {
		hashedNewPassword = user.Password
	}

	if username == "" {
		username = user.Username
	}

	_, err = DB.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", username, hashedNewPassword, user.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
