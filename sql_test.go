package golang_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := DBConnection()
	defer db.Close()

	ctx := context.Background()
	sql_query := "INSERT INTO customers (name, age, city) VALUES ('Tejo', '18', 'Semarang')"
	_, err := db.ExecContext(ctx, sql_query)
	// Gunakan ExecContext untuk eksekusi query yang mengeksekusi database, seperti, insert, update, delete
	if err != nil {
		panic(err)
	}

	fmt.Println("Data Inserted Succesfully")
}

func TestQuerySql(t *testing.T) {
	db := DBConnection()
	defer db.Close()

	ctx := context.Background()
	sql_query := "SELECT * FROM customers"
	rows, err := db.QueryContext(ctx, sql_query)
	// Gunakan QueryContext untuk eksekusi query yang menampilkan isi database, seperti, select
	if err != nil {
		panic(err)
	}

	defer rows.Close()
}

func TestQuerySqlComplex(*testing.T) {
	db := DBConnection()
	defer db.Close()

	ctx := context.Background()
	sql_query := "SELECT email, name, age, city, balance, rating, birth_date, verified, created_at FROM customers"
	rows, err := db.QueryContext(ctx, sql_query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// rows.Next adalah untuk memilih data dari awal hingga akhir
	for rows.Next(){
		// golang tidak support untuk nilai NULL dari database, sehingga harus merubah tipe data dari column yang bernilai NULL
		var name, city string
		var email sql.NullString // jika nilai NULL dari database dengan tipe data string
		var age, balance int
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime // jika nilai NULL dari database dengan tipe data time 
		var verified bool

		//rows.Scan adalah untuk mencari apa data yang dibutuhkan
		err := rows.Scan(&email, &name, &age, &city, &balance, &rating, &birth_date, &verified, &created_at)
		if err != nil{
			panic(err)
		}
		fmt.Println("=====================================")
		// Dari sql.NullTipeData akan mengembalikan VALID nilai TRUE atau FALSE, sehingga bisa menggunakan IF untuk menampilkan datanya
		// Jika balikan VALID adalah FALSE, maka data tidak akan ditampilkan, jika VALID adalah TRUE data akan ditampilkan
		if email.Valid{
			fmt.Println("email:" , email.String)
		} else {
			fmt.Println("email: Not Valid")
		}
		fmt.Println("name:" , name)
		fmt.Println("age:" , age)
		fmt.Println("city:" , city)
		fmt.Println("balance:" , balance)
		fmt.Println("rating:" , rating, "/ 5")
		if birth_date.Valid{
			fmt.Println("birth_date:" , birth_date.Time)
		} else {
			fmt.Println("birth_date: Not Valid")
		}
		fmt.Println("verified:" , verified)
		fmt.Println("created_at:" , created_at)
	}
}