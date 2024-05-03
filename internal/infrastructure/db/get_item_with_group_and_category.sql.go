// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_item_with_group_and_category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getItemWithGroupAndCategory = `-- name: GetItemWithGroupAndCategory :one
SELECT items.item_id, items.name, items.description, items.category_id, items.group_id, items.location_id, items.is_stored, items."createdAt", items."updatedAt",
       categories.name        AS category_name,
       categories.description AS category_description,
       groups.name            AS group_name,
       groups.description     AS group_description

FROM items
         LEFT JOIN categories ON items.category_id = categories.category_id
         LEFT JOIN groups ON items.group_id = groups.group_id
WHERE items.item_id = $1
`

type GetItemWithGroupAndCategoryRow struct {
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
	GroupName           pgtype.Text
	GroupDescription    pgtype.Text
}

// noinspection SqlResolveForFile
func (q *Queries) GetItemWithGroupAndCategory(ctx context.Context, itemID int32) (GetItemWithGroupAndCategoryRow, error) {
	row := q.db.QueryRow(ctx, getItemWithGroupAndCategory, itemID)
	var i GetItemWithGroupAndCategoryRow
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
		&i.GroupName,
		&i.GroupDescription,
	)
	return i, err
}
