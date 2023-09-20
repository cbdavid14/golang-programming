package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang-programming/dependency-injection/configuration"
	"log"
)

func CreateSqliteConnection(ctx context.Context, config *configuration.Configuration) (*sqlx.DB, error) {
	log.Println("Creating sqlite connection")
	source := fmt.Sprintf("./dependency-injection/db/%s.db", config.DB.Name)
	db, err := sqlx.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Sqlite connection created")
	return db, nil
}
