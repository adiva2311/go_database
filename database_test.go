package golang_mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_mysql")
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to Database Successfully")
	defer db.Close()
}
