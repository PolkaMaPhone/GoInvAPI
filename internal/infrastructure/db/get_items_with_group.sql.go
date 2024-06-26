// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_items_with_group.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllItemsWithGroups = `-- name: GetAllItemsWithGroups :many
SELECT items.item_id, items.name, items.description, items.category_id, items.group_id, items.location_id, items.is_stored, items."createdAt", items."updatedAt",
       groups.name        AS group_name,
       groups.description AS group_description
FROM items
         LEFT JOIN groups ON items.group_id = groups.group_id
ORDER BY items.item_id
`

type GetAllItemsWithGroupsRow struct {
	ItemID           int32
	Name             string
	Description      pgtype.Text
	CategoryID       pgtype.Int4
	GroupID          pgtype.Int4
	LocationID       pgtype.Int4
	IsStored         pgtype.Bool
	CreatedAt        pgtype.Timestamptz
	UpdatedAt        pgtype.Timestamptz
	GroupName        pgtype.Text
	GroupDescription pgtype.Text
}

// noinspection SqlResolveForFile
func (q *Queries) GetAllItemsWithGroups(ctx context.Context) ([]GetAllItemsWithGroupsRow, error) {
	rows, err := q.db.Query(ctx, getAllItemsWithGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllItemsWithGroupsRow
	for rows.Next() {
		var i GetAllItemsWithGroupsRow
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
			&i.GroupName,
			&i.GroupDescription,
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
