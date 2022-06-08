package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://postgres:@localhost:5432/?sslmode=disable"),
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithUser(viper.GetString("USER")),
		pgdriver.WithPassword(viper.GetString("PASSWORD")),
		pgdriver.WithTimeout(5*time.Second),
	),
	)

	db := bun.NewDB(sqldb, pgdialect.New())
	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while pinging postgres")
	}
	return db, nil
}

func Connection() (*bun.DB, error) {
	con, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error while connecting to db")
	}
	return con, nil
}
