package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gui-laranjeira/todo-api/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = db.Ping()
	return db, err
}
