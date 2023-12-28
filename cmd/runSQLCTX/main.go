package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Jhon-Henkel/go-lang-full-cycle-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDb struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDb(dbConn *sql.DB) *CourseDb {
	return &CourseDb{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDb) callTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := db.New(tx)
	err = fn(query)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDb) CreateCourseAndCategory(ctx context.Context, argsCourse CourseParams, argsCategory CategoryParams) error {
	err := c.callTx(ctx, func(queries *db.Queries) error {
		var err error
		err = queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = queries.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	//queries := NewCourseDb(dbConn)

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Curso de Go",
		Description: sql.NullString{"Descrição do curso de Go", true},
		Price:       10.95,
	}

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{"Categoria do curso de Go", true},
	}

	courseDb := NewCourseDb(dbConn)
	err = courseDb.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	if err != nil {
		panic(err)
	}
}
