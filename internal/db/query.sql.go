// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?)
`

type CreateCategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name, arg.Description)
	return err
}

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price) VALUES (?, ?, ?, ?, ?)
`

type CreateCourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  string
	Price       float64
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.Price,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?
`

func (q *Queries) DeleteCategory(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, description FROM categories WHERE id = ?
`

func (q *Queries) GetCategory(ctx context.Context, id string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, description FROM categories
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?
`

type UpdateCategoryParams struct {
	Name        string
	Description sql.NullString
	ID          string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.Name, arg.Description, arg.ID)
	return err
}
