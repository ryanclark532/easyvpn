package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

type Log struct {
	bun.BaseModel `bun:"table:logs,alias:l"`
	time          time.Time `bun:",notnull"`
	text          string    `bun:",notnull"`
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

/*
func SetupTestData() error {
	password, err := bcrypt.GenerateFromPassword([]byte("test1"), 10)
	if err != nil {
		return err
	}
	user := user.User{
		Name:           "test1",
		Username:       "test1",
		Password:       string(password),
		PasswordExpiry: time.Now().Add(time.Hour * 24),
		IsAdmin:        true,
		Enabled:        true,
		Roles:          "Server Status,User Management",
	}

	group := groups_dtos.Group{
		Name:        "group1",
		IsAdmin:     true,
		Enabled:     true,
		Roles:       "Server Status,User Management,Settings",
		MemberCount: 1,
	}

	groupMembership := groups_dtos.GroupMembership{
		GroupID: 1,
		UserID:  1,
	}

	settings := &settings_dtos.Settings{
		Version:         0,
		Latest:          true,
		AllowChangePW:   true,
		EnforceStrongPW: true,
		MaxAuthAttempts: 3,
		LockoutTimeout:  10000,
		WebServerPort:   8080,
		IPAddress:       "192.168.86.10",
		MaxConnections:  5,
	}

	_, err = DB.NewInsert().Model(&user).Exec(context.Background())
	if err != nil {
		return err
	}
	_, err = DB.NewInsert().Model(&group).Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = DB.NewInsert().Model(&groupMembership).Exec(context.Background())
	if err != nil {
		return err
	}
	_, err = DB.NewInsert().Model(settings).Exec(context.Background())
	return err
}
*/
