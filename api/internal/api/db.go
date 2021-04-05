package api

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"os"
)

var models = []interface{}{
	(*User)(nil),
}

func init() {
	db := getDatabaseConnection()
	defer db.Close()
	createSchema(db)
}

func createSchema(db *pg.DB) error {
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func getDatabaseConnection() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})
}
