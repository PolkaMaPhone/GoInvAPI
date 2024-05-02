// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_items_with_category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllItemsWithCategories = `-- name: GetAllItemsWithCategories :many
SELECT items.item_id, items.name, items.description, items.category_id, items.group_id, items.location_id, items.is_stored, items."createdAt", items."updatedAt",
       categories.name AS category_name,
       categories.description AS category_description
FROM items
         JOIN categories ON items.category_id = categories.category_id
ORDER BY items.item_id
`

type GetAllItemsWithCategoriesRow struct {
	ItemID              int32
	Name                string
	Description         pgtype.Text
	CategoryID          pgtype.Int4
	GroupID             pgtype.Int4
	LocationID          pgtype.Int4
	IsStored            pgtype.Bool
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	CategoryName        string
	CategoryDescription pgtype.Text
}

// noinspection SqlResolveForFile
func (q *Queries) GetAllItemsWithCategories(ctx context.Context) ([]GetAllItemsWithCategoriesRow, error) {
	rows, err := q.db.Query(ctx, getAllItemsWithCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllItemsWithCategoriesRow
	for rows.Next() {
		var i GetAllItemsWithCategoriesRow
		if err := rows.Scan(
			&i.ItemID,
			&i.Name,
			&i.Description,
			&i.CategoryID,
			&i.GroupID,
			&i.LocationID,
			&i.IsStored,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryName,
			&i.CategoryDescription,
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
