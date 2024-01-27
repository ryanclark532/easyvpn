package database

import (
	"context"
	"database/sql"
	"easyvpn/src/groups/groups_dtos"
	"easyvpn/src/settings/settings_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"golang.org/x/crypto/bcrypt"
)

var DB *bun.DB

type Log struct {
	bun.BaseModel `bun:"table:logs,alias:l"`
	time          time.Time `bun:",notnull"`
	text          string    `bun:",notnull"`
}

func Test() error {
	_, err := os.Stat("database.db")
	if err == nil {
		err = TestGetDB()
		if err != nil {
			return err
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create("database.db")
	if err != nil {
		return err
	}
	file.Close()

	err = TestGetDB()
	if err != nil {
		return err
	}
	DB.NewCreateTable().Model((*user_dtos.User)(nil)).Exec(context.Background())
	DB.NewCreateTable().Model((*groups_dtos.Group)(nil)).Exec(context.Background())
	DB.NewCreateTable().Model((*groups_dtos.GroupMembership)(nil)).Exec(context.Background())
	DB.NewCreateTable().Model((*settings_dtos.Settings)(nil)).Exec(context.Background())
	DB.NewCreateTable().Model((*Log)(nil)).Exec(context.Background())
	err = SetupTestData()
	if err != nil {
		return err
	}

	return nil

}

func TestGetDB() error {
	sqldb, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	DB = db
	return nil
}

func SetupTestData() error {
	password, err := bcrypt.GenerateFromPassword([]byte("test1"), 10)
	if err != nil {
		return err
	}
	user := user_dtos.User{
		Name:           "test1",
		Username:       "test1",
		Password:       string(password),
		PasswordExpiry: time.Now().Add(time.Hour * 24),
		IsAdmin:        true,
		Enabled:        true,
		Roles:          "vpn, usermgt, settings",
	}

	group := groups_dtos.Group{
		Name:        "group1",
		IsAdmin:     true,
		Enabled:     true,
		Roles:       "vpn, usermgt, settings",
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
