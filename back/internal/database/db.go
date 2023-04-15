package database

import (
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

func New(dsn string) (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("postgres", "postgres://"+dsn)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(2 * time.Hour)

	return DB, nil
}
