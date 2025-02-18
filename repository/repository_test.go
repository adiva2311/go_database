package repository

import (
	"context"
	"fmt"
	"golang_mysql"
	"golang_mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	reviewRepo := NewReviewsRepository(golang_mysql.DBConnection())
	
	ctx := context.Background()
	review := entity.Reviews{
		Email: "kurakura@gmail.com",
		Comment: "Saya suka makan",
	}

	insert, err := reviewRepo.InsertData(ctx, review)
	if err != nil{
		panic(err)
	}

	fmt.Println(insert)
}

func TestUpdate(t *testing.T){
	reviewRepo := NewReviewsRepository(golang_mysql.DBConnection())

	ctx := context.Background()
	review := entity.Reviews{
		Email: "sukawedi@gmail.com",
		Comment: "Saya suka denger podcast",
	}

	update, err := reviewRepo.UpdateData(ctx, review, 5)
	if err != nil{
		panic(err)
	}

	fmt.Println(update)
}

func TestDelete(t *testing.T){
	reviewRepo := NewReviewsRepository(golang_mysql.DBConnection())

	delete, err := reviewRepo.DeleteData(context.Background(), 1)
	if err != nil{
		panic(err)
	}

	fmt.Println(delete)
}

func TestFindAll(t *testing.T) {
	reviewRepo := NewReviewsRepository(golang_mysql.DBConnection())

	ctx := context.Background()

	findAll, err := reviewRepo.FindAll(ctx)
	if err != nil{
		panic(err)
	}

	fmt.Println(findAll)
}

func TestFindById(t *testing.T) {
	reviewRepo := NewReviewsRepository(golang_mysql.DBConnection())

	ctx := context.Background()

	findById, err := reviewRepo.FindById(ctx, 6)
	if err != nil{
		panic(err)
	}

	fmt.Println(findById)
}