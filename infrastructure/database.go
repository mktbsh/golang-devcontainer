package infrastructure

import (
	"context"
	"database/sql"

	"github.com/mktbsh/golang-devcontainer/config"
	"github.com/mktbsh/golang-devcontainer/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

func init() {
	c := config.Config.Database
	sqldb, err := sql.Open(sqliteshim.ShimName, c.Path)
	if err != nil {
		panic(err)
	}

	DB = bun.NewDB(sqldb, sqlitedialect.New())

	setDebugOptions()

	DB.NewCreateTable().Model((*model.User)(nil)).IfNotExists().Exec(context.Background())
}

func setDebugOptions() {
	if config.Config.IsProduction() {
		return
	}

	DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
}
