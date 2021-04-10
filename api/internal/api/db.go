package api

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/zap"
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
	log, _ := zap.NewProduction()
	defer log.Sync()
	user := os.Getenv("POSTGRES_USER")
	if len(user) == 0 {
		panic("ERROR:  POSTGRES_USER not set")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if len(password) == 0 {
		panic("ERROR:  POSTGRES_PASSWORD not set")
	}
	database := os.Getenv("POSTGRES_DB")
	if len(database) == 0 {
		panic("ERROR:  POSTGRES_DATABASE not set")
	}
	log.Info(fmt.Sprintf("user=%s, password=\"%s\", db=%s", user, password, database))
	return pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
}
