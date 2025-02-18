package golang_mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_mysql?parseTime=true")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database Successfully")
	
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	
	return db
}