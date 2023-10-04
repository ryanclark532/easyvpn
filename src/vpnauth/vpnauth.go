package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var username string = os.Args[1]
var password string = os.Args[2]

type user struct {
	username string
	password string
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	user, err := getUser(db)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if user.password == password {
		fmt.Print("Passwords dont match")
		os.Exit(1)
	}

	os.Exit(0)
}

func getUser(db *sql.DB) (*user, error) {
	var user user
	fmt.Println(username)
	query := fmt.Sprintf("SELECT username, password FROM Users WHERE username ='%s'", username)
	err := db.QueryRow(query).Scan(&user.username, &user.password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
