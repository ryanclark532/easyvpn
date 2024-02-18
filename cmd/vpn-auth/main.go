package main

import (
	"context"
	"easyvpn/internal/database"
	"easyvpn/internal/login"
	"easyvpn/internal/user"
	"fmt"
	"os"
	"strings"
	"unicode"

	_ "github.com/mattn/go-sqlite3"
)

func removeWhitespace(input string) string {
	var result []rune

	for _, char := range input {
		if !unicode.IsSpace(char) {
			result = append(result, char)
		}
	}

	return string(result)
}

func GetUser(username string) (*user.User, error) {
	var user = new(user.User)
	err := database.DB.NewSelect().Model(user).Where("username = ?", username).Limit(1).Scan(context.Background(), user)
	return user, err
}

func main() {
	fmt.Println(os.Args)

	err := database.GetDB()
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

	authed, _, err := login.AuthUser(username, password)
	if err != nil {
		_ = os.WriteFile("output.txt", []byte(err.Error()), 0644)
		os.Exit(1)
	}
	if !authed {
		os.Exit(1)
	}

	os.Exit(0)
}
