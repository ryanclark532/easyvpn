package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

func Test() error {
	_, err := os.Stat("test.db")
	if err == nil {
		err = TestGetDB()
		if err != nil {
			return err
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create("test.db")
	if err != nil {
		return err
	}
	file.Close()

	err = TestGetDB()
	if err != nil {
		return err
	}

	return nil

}

func TestGetDB() error {
	sqldb, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	DB = db
	return nil
}

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

	err = setupTables(db)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = db.Exec(`INSERT INTO Users (username, name, password, is_admin, enabled, password_expiry)
                      VALUES (?, ?, ?, ?, ?, ?);`,
		"test",
		"Dummy User",
		"test",
		true,
		true,
		time.Now().AddDate(0, 0, 30).Format(time.DateTime),
	)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO Groups (name,member_count, enabled,is_admin)
                      VALUES (?, ?, ?, ?);`,
		"test",
		0,
		true,
		true,
	)
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

func setupTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		is_admin BOOLEAN,
	    enabled BOOLEAN,
		password_expiry DATE NOT NULL
	);`)

	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Groups (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		member_count INTEGER NOT NULL,
		enabled BOOLEAN NOT NULL,
		is_admin BOOLEAN NOT NULL
	);`)

	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS GroupMembership (
		id INTEGER PRIMARY KEY,
		groupId INTEGER NOT NULL,
		FOREIGN KEY (groupId) REFERENCES Groups(id)
		userId INTEGER NOT NULL,
		FOREIGN KEY(userId) REFERENCES Users(id)
	);`)
	if err != nil {
		return err
	}

	return nil

}
