package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"

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
	err = db.QueryRow("SELECT id, email, password, username, root_folder FROM users WHERE id=?", userID).Scan(&u.ID, &u.Email, &u.Password, &u.Email, &rootID)
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
	var user User
	user.Email = r.PostFormValue("email")
	user.Password = r.PostFormValue("password")

	// open
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// run it
	row := db.QueryRow("SELECT id,password FROM users WHERE email = ?", user.Email)

	var dbPass string
	err = row.Scan(&user.ID, &dbPass)
	if err != nil {
		http.Error(w, "bad username or password", 400)
	}

	sltpwd := append([]byte(user.Password), pwdsalt...)
	dbBytepass := []byte(dbPass)

	err = bcrypt.CompareHashAndPassword(dbBytepass, sltpwd)

	if err != nil {
		http.Error(w, "bad username or password", 400)
	}
	user.CreateSession(w)
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	/*
				cookie, cookieErr := r.Cookie("session_id")
				if cookieErr != nil {
					// the user did not give us a cookie to logout with
					// they probably typed in the exact url for logout requests
					// TODO 404 them
					return
				}

				// Salt the session and generate checksum

				//saltedChecksum := bytestream the cookie.value
				// then concat the salt from session.go
				// TODO a function that literally does this in session.go
				// then generates either bytestream or string
				// then we can hand it off to sql

				// find session cookie in database and clear if there
				db, sqlErr := sql.Open("sqlite3", DatabaseFile)
				if sqlErr != nil {
					return sqlErr
				}
				defer db.Close()
				row, dbErr := db.Query("DELETE * FROM sessions WHERE checksum = '" + saltedChecksum + "'")
		=======
			// find session cookie in database and clear if there
			//db, err := sql.Open("sqlite3", DatabaseFile)
			//if err != nil {
			//	return err
			//	http.Error(w, err.Error(), 500)
			//	return
			//}
			//defer db.Close()
			//err = db.Exec("DELETE FROM sessions WHERE checksum = ?", saltedChecksum)
			//if err != nil && err != sql.ErrNoRows {
			//	http.Error(w, err.Error(), 500)
			//	return
			//}
		>>>>>>> Stashed changes

				// tell user to clear cookie regardless
				cookie.MaxAge = -1
				http.SetCookie(w, cookie)
	*/
}
