package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel  `bun:"table:users,alias:u"`
	ID             uint      `bun:",pk,autoincrement" json:"id"`
	Name           string    `bun:",notnull" json:"name"`
	Username       string    `bun:",notnull" json:"username"`
	Password       string    `bun:",notnull" json:"password"`
	IsAdmin        bool      `json:"is_admin"`
	Enabled        bool      `json:"enabled"`
	PasswordExpiry time.Time `json:"password_expiry"`
}

var DB *bun.DB

func removeWhitespace(input string) string {
	var result []rune

	for _, char := range input {
		if !unicode.IsSpace(char) {
			result = append(result, char)
		}
	}

	return string(result)
}
func GetDB() error {
	sqldb, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	DB = db
	return nil
}

func GetUser(username string) (*User, error) {
	var user = new(User)
	err := DB.NewSelect().Model(user).Where("username = ?", username).Limit(1).Scan(context.Background(), user)
	return user, err
}

func AuthUser(username string, password string) error {
	user, err := GetUser(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println(os.Args)

	err := GetDB()
	if err != nil {
		_ = os.WriteFile("output.txt", []byte(err.Error()), 0644)
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		_ = os.WriteFile("output.txt", []byte(err.Error()), 0644)
		os.Exit(1)
	}

	s := strings.Split(string(content), "\n")
	username := removeWhitespace(s[0])
	password := removeWhitespace(s[1])
	if username == "" || password == "" {
		os.Exit(1)
	}

	err = AuthUser(username, password)
	if err != nil {
		_ = os.WriteFile("output.txt", []byte(err.Error()), 0644)
		os.Exit(1)
	}

	os.Exit(0)
}
