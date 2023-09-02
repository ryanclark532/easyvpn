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
		ID INTEGER PRIMARY KEY,
		Username VARCHAR(255) UNIQUE NOT NULL,
		Name VARCHAR(255) NOT NULL,
		Password VARCHAR(255) NOT NULL,
		IsAdmin BOOLEAN NOT NULL,
	    Enabled BOOLEAN NOT NULL

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
