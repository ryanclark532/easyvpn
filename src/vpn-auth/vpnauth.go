package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"golang.org/x/crypto/bcrypt"
)

var DB *bun.DB

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint
	Name           string
	Username       string
	Password       string
	IsAdmin        bool
	Enabled        bool
	PasswordExpiry time.Time
}

func getDB() (*bun.DB, error) {
	sqldb, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	return db, nil
}

func getUser(username string) (*User, error) {
	user := new(User)

	db, err := getDB()
	if err != nil {
		return nil, err
	}

	err = db.NewSelect().Model(user).Where("username = ?", username).Limit(1).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func authUser(username string, password string) (bool, error) {
	user, err := getUser(username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func main() {
	var username string = os.Args[1]
	var password string = os.Args[2]
	authed, err := authUser(username, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if authed {
		fmt.Println("Sucess")
		os.Exit(0)
	}
	fmt.Println("Failed")
	os.Exit(1)
}
