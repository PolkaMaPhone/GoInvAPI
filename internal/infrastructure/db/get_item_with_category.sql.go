// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_item_with_category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getItemWithCategory = `-- name: GetItemWithCategory :one
SELECT items.item_id, items.name, items.description, items.category_id, items.group_id, items.location_id, items.is_stored, items."createdAt", items."updatedAt",
       categories.name        AS category_name,
       categories.description AS category_description
FROM items
         LEFT JOIN categories ON items.category_id = categories.category_id
WHERE item_id = $1
`

type GetItemWithCategoryRow struct {
	ItemID              int32
	Name                string
	Description         pgtype.Text
	CategoryID          pgtype.Int4
	GroupID             pgtype.Int4
	LocationID          pgtype.Int4
	IsStored            pgtype.Bool
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	CategoryName        pgtype.Text
	CategoryDescription pgtype.Text
}

// noinspection SqlResolveForFile
func (q *Queries) GetItemWithCategory(ctx context.Context, itemID int32) (GetItemWithCategoryRow, error) {
	row := q.db.QueryRow(ctx, getItemWithCategory, itemID)
	var i GetItemWithCategoryRow
	err := row.Scan(
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
	)
	return i, err
}
