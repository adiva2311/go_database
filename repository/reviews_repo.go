package repository

import (
	"context"
	"golang_mysql/entity"
)


type ReviewsRepository interface {
	InsertData(ctx context.Context, review entity.Reviews) (entity.Reviews, error)
	UpdateData(ctx context.Context, review entity.Reviews,  Id int) (entity.Reviews, error)
	DeleteData(ctx context.Context, Id int) (entity.Reviews, error)
    FindById(ctx context.Context, Id int) (entity.Reviews, error)
    FindAll(ctx context.Context) ([]entity.Reviews, error)
}