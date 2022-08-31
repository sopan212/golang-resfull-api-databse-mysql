package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_basic_sql")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return db, nil

}
