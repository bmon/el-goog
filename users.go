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
	ID       int
	Email    string
	Password string
	Username string
}

func (u *User) Insert() error {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := "INSERT INTO users(id, email, password, username) values(?,?,?,?)"
	res, err := db.Exec(sqlStmt, nil, u.Email, u.Password, u.Username)
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

func UserCreate(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	username := r.PostFormValue("username")

	var isEmail = regexp.MustCompile(`^.+\@.+\..+$`)

	if isEmail.MatchString(email) {

                sltpwd := append([]byte(password), pwdsalt...)
                hshpwd, _ := bcrypt.GenerateFromPassword(sltpwd, 10) //salting and hashing the password

                hashedPassword := string(hshpwd[:])

		user := &User{-1, email, hashedPassword, username}
		err := user.Insert()

		if err != nil {
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
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

        db, err := sql.Open("sqlite3", DatabaseFile)
        row, err :=db.Query("SELECT password FROM users WHERE email = '"+email+"'")
        if err != nil {
             fmt.Println(err)
        }

        var recPassword []byte

        sltpwd := append([]byte(password), pwdsalt...)
        
        err = row.Scan(&recPassword)

        if err == nil {
            fmt.Println("error scanning row")
        }
       
        err = bcrypt.CompareHashAndPassword(recPassword, sltpwd)


        if err == nil {
            fmt.Println("match!")
        } else {
            fmt.Println("try again lel")
        }
        
        //check that email address exists
        //create struct instance
        //check password
        //if yes - instance.CreateSession
}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	//cookie := 
	

}
