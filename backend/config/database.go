package config

import (
	"database/sql"
	"fmt"
	"kel15/services"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLLiteDB() (*services.SQLLite3, error) {
	name := os.Getenv("SQLITE3_NAME")
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", name))
	if err != nil {
		return nil, err
	}

	return &services.SQLLite3{
		DB: db,
	}, nil

}
