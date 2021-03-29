package api

import (
	"github.com/go-pg/pg/v10"
	"os"
)

func getDatabaseConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})
}
