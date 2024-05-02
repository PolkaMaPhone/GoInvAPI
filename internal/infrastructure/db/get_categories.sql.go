// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_categories.sql

package db

import (
	"context"
)

const getAllCategories = `-- name: GetAllCategories :many

SELECT category_id, name, description, "createdAt", "updatedAt"
FROM categories
ORDER BY category_id
`

// noinspection SqlResolveForFile
func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}