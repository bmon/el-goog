package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var salt []byte = []byte(getEnv("COOKIE_SALT", "something-very-secret"))
var expiryDays int = 30
var keySize int = 200

type Session struct {
	UserID   int
	Expires  time.Time
	Checksum [sha256.Size]byte
}

func (s *Session) insert() {
	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := fmt.Sprintf("insert into sessions values (NULL, %d, %d, '%x')", s.UserID, s.Expires.Unix(), s.Checksum)
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}
}

func (u *User) CreateSession(w http.ResponseWriter) {
	sessionID := make([]byte, keySize) // make empty byte slice
	_, err := rand.Read(sessionID)     // fill it with random data
	if err != nil {
		fmt.Println(err)
	}

	saltedSessionID := append(sessionID, salt...)

	newSession := &Session{
		u.ID,
		time.Now().AddDate(0, 0, expiryDays),
		sha256.Sum256(saltedSessionID),
	}

	newSession.insert()
	fmt.Println(base64.StdEncoding.EncodeToString(sessionID))
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  base64.StdEncoding.EncodeToString(sessionID),
		MaxAge: expiryDays * 24 * 60 * 60,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "root_id",
		Value:  fmt.Sprintf("%d", u.RootFolder.ID),
		MaxAge: expiryDays * 24 * 60 * 60,
	})
}

func GetRequestUser(r *http.Request) *User {
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		// the cookie session_id was not found, so return a nil user.
		return nil
	}

	sessionBytes, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil || len(sessionBytes) != keySize {
		// the cookie value was bad
		return nil
	}

	saltedSessionID := append(sessionBytes, salt...)

	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var userID int

	stmt := fmt.Sprintf("select user_id from sessions where checksum='%x' and expires > %d", sha256.Sum256(saltedSessionID), time.Now().Unix())
	row := db.QueryRow(stmt)
	switch err := row.Scan(&userID); err {
	case sql.ErrNoRows:
		// no such session id exists, or it's expired.
		return nil
	case nil:
		// the user session exists and is valid!
		user, err := UserSelectByID(userID)
		if err != nil {
			return nil
		}
		return user
	default:
		// something else horrible!
		fmt.Println("ERROR retrieving user sesion:", err)
		return nil
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
