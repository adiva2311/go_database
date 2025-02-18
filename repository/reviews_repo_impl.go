package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang_mysql/entity"
	"strconv"
)

type reviewsRepositoryImpl struct {
	DB *sql.DB
}

// DeleteData implements ReviewsRepository.
func (r *reviewsRepositoryImpl) DeleteData(ctx context.Context, Id int) (entity.Reviews, error) {
	scriptSql := "DELETE FROM reviews WHERE id = ?"
	result, err := r.DB.ExecContext(ctx, scriptSql, Id)
	hasil := entity.Reviews{}
	if err != nil{
		return hasil, err
	}

	affectId, err := result.RowsAffected()
	if err != nil {
		return hasil, err
	}

	//fmt.Println("Rows Affected :", affectId)
	if affectId > 0{
		fmt.Println("Success Delete Data with ID :", Id)
	} else {
		fmt.Println("No Data with ID :", Id)
	}
	return hasil, nil
}

// FindAll implements ReviewsRepository.
func (r *reviewsRepositoryImpl) FindAll(ctx context.Context) ([]entity.Reviews, error) {
	scriptSql := "SELECT id, email, comment FROM reviews"
	rows, err := r.DB.QueryContext(ctx, scriptSql)
	if err != nil {
		return nil, err
	}

	var reviews []entity.Reviews
	for rows.Next() {
		review := entity.Reviews{}
		rows.Scan(&review.Id, &review.Email, &review.Comment)
		reviews = append(reviews, review)
	}
	defer rows.Close()

	return reviews, nil
}

// FindById implements ReviewsRepository.
func (r *reviewsRepositoryImpl) FindById(ctx context.Context, Id int) (entity.Reviews, error) {
	scriptSql := "SELECT id, email, comment FROM reviews WHERE id = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, scriptSql, Id)
	review := entity.Reviews{}
	if err != nil {
		return review, err
	}

	if rows.Next() {
		err := rows.Scan(&review.Id, &review.Email, &review.Comment)
		if err != nil {
			return review, err
		}
	} else {
		return review, errors.New("Can't Find ID with : " + strconv.Itoa(Id))
	}
	defer rows.Close()

	return review, nil
}

// UpdateData implements ReviewsRepository.
func (r *reviewsRepositoryImpl) UpdateData(ctx context.Context, review entity.Reviews, Id int) (entity.Reviews, error) {
	scriptSql := "UPDATE reviews SET email = ?, comment = ? WHERE id = ?"
	result, err := r.DB.ExecContext(ctx, scriptSql, review.Email, review.Comment, Id)
	if err != nil {
		return review, err
	}

	affectId, err := result.RowsAffected()
	if err != nil {
		return review, err
	}

	fmt.Println("Rows Affected :", affectId)
	review.Id = int32(Id)
	return review, nil
}

// InsertData implements ReviewsRepository.
func (r *reviewsRepositoryImpl) InsertData(ctx context.Context, review entity.Reviews) (entity.Reviews, error) {
	scriptSql := "INSERT INTO reviews (email, comment) VALUES (?,?)"
	result, err := r.DB.ExecContext(ctx, scriptSql, review.Email, review.Comment)
	if err != nil {
		return review, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return review, err
	}

	fmt.Println("Success Insert Comment with ID :", insertedId)
	review.Id = int32(insertedId)
	return review, nil
}

func NewReviewsRepository(db *sql.DB) ReviewsRepository {
	return &reviewsRepositoryImpl{DB: db}
}
