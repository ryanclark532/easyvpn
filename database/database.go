package database

import (
	"context"
	"database/sql"
	"easyvpn/groups/groups_dtos"
	user_dtos "easyvpn/user/user-dtos"
	"os"
	"time"

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
    user:= user_dtos.User{
        Name: "test1",
        Username: "test1",
        Password: "test1",
        PasswordExpiry: time.Now().Add(time.Hour *24),
        IsAdmin: true,
        Enabled: true,
    }

    group := groups_dtos.Group{
       Name: "group1",
       IsAdmin: true,
       Enabled: true,
       MemberCount: 1,
    }

    groupMembership := groups_dtos.GroupMembership{
       GroupID: 1,
       UserID: 1,
    }

    _, err := DB.NewInsert().Model(&user).Exec(context.Background())
    if err != nil{
        return err
    }
    _, err = DB.NewInsert().Model(&group).Exec(context.Background())
    if err != nil{
        return err
    }

    _, err = DB.NewInsert().Model(&groupMembership).Exec(context.Background())
    return err
}
