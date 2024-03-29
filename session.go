package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
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
	sqlStmt := fmt.Sprintf("insert into sessions values (NULL, %d, %d, '%x')", s.UserID, s.Expires.Unix(), s.Checksum)
	_, err := DB.Exec(sqlStmt)
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
	http.SetCookie(w, &http.Cookie{
		Name:   "user_id",
		Value:  fmt.Sprintf("%d", u.ID),
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

	var userID int

	stmt := fmt.Sprintf("select user_id from sessions where checksum='%x' and expires > %d", sha256.Sum256(saltedSessionID), time.Now().Unix())
	row := DB.QueryRow(stmt)
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
func DeleteRequestSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		// the cookie session_id was not found, so return a nil user.
		return
	}

	sessionBytes, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil || len(sessionBytes) != keySize {
		// the cookie value was bad
		return
	}

	fmt.Println("DELET THIS")
	// tell the browser to delet this cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})

	saltedSessionID := append(sessionBytes, salt...)

	DB.Exec("delete from sessions where checksum=?", fmt.Sprintf("%x", sha256.Sum256(saltedSessionID)))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
