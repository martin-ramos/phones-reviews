package database

import (
	"database/sql"
	"fmt"
)

type MySqlClient struct {
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		fmt.Errorf("Cannot create db client: %s", err.Error())
		panic("error")
	}

	return db
}
