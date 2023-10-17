package database

import (
	"context"
	"database/sql"
	"easyvpn/src/groups/groups_dtos"
	user_dtos "easyvpn/src/user/user-dtos"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

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
