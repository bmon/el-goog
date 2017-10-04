package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
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

func CreateSession(userID int) string {
	sessionID := make([]byte, keySize) // make empty byte slice
	_, err := rand.Read(sessionID)     // fill it with random data
	if err != nil {
		fmt.Println(err)
	}

	saltedSessionID := append(sessionID, salt...)

	newSession := &Session{
		userID,
		time.Now().AddDate(0, 0, expiryDays),
		sha256.Sum256(saltedSessionID),
	}

	newSession.insert()

	return base64.StdEncoding.EncodeToString(sessionID)
}

func GetSessionUserID(sessionID string) (int, error) {
	sessionBytes, err := base64.StdEncoding.DecodeString(sessionID)
	if err != nil {
		return -1, err
	}
	if len(sessionBytes) != keySize {
		return -1, errors.New("Bad sessionID given")
	}
	saltedSessionID := append(sessionBytes, salt...)

	db, err := sql.Open("sqlite3", DatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user_id int

	fmt.Printf("%x\n", sha256.Sum256(saltedSessionID))

	stmt := fmt.Sprintf("select user_id from sessions where checksum='%x' and expires > %d", sha256.Sum256(saltedSessionID), time.Now().Unix())
	row := db.QueryRow(stmt)
	switch err := row.Scan(&user_id); err {
	case sql.ErrNoRows:
		return -1, err
	case nil:
		return user_id, nil
	default:
		log.Fatal(err)
	}
	return -1, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
