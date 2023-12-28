package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Jhon-Henkel/go-lang-full-cycle-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	ValidateError(err)
	defer dbConn.Close()

	queries := db.New(dbConn)
	id := uuid.New().String()

	// Create
	fmt.Println("Start Create")
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Category 1",
		Description: sql.NullString{String: "Description 1", Valid: true},
	})
	ValidateError(err)
	ListCategories(queries, ctx)
	fmt.Println("End Create")

	// Update
	fmt.Println("--------------------")
	fmt.Println("Start Update")
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          id,
		Name:        "Category 1 Updated",
		Description: sql.NullString{String: "Description 1 Updated", Valid: true},
	})
	ValidateError(err)
	ListCategories(queries, ctx)
	fmt.Println("End Update")

	// Delete
	fmt.Println("--------------------")
	fmt.Println("Start Delete")
	err = queries.DeleteCategory(ctx, id)
	ValidateError(err)
	ListCategories(queries, ctx)
	fmt.Println("End Delete")
}

func ListCategories(queries *db.Queries, ctx context.Context) {
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
}

func ValidateError(err error) {
	if err != nil {
		panic(err)
	}
}
