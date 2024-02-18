package logging

import (
	"context"
	"easyvpn/internal/database"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Log struct {
	bun.BaseModel `bun:"table:logs,alias:l"`
	time          time.Time `bun:",notnull"`
	text          string    `bun:",notnull"`
}

func HandleError(err error, errorFunction string) {
	fmt.Println(err.Error() + " " + errorFunction)
	e := &Log{
		time: time.Now(),
		text: err.Error() + " " + errorFunction,
	}
	database.DB.NewInsert().Model(e).Exec(context.Background())
}
