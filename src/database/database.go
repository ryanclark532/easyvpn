package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitializeDatabase() error {
	_, err := os.Stat("database.db")
	if err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create("database.db")
	if err != nil {
		return err
	}
	file.Close()

	db, err := GetDB()
	if err != nil {
		return err
	}

	// TODO extract this to a method
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		is_admin BOOLEAN NOT NULL,
	    enabled BOOLEAN NOT NULL

	);`)
	if err != nil {
		return err
	}

	return nil
}

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
