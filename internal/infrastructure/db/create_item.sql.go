package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
)

const createItem = `-- name: CreateItem :execlastid
INSERT INTO items (name, description, category_id, group_id, location_id, is_stored, "createdAt", "updatedAt")
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING item_id
`

type CreateItemParams struct {
	Name        string
	Description pgtype.Text
	CategoryID  pgtype.Int4
	GroupID     pgtype.Int4
	LocationID  pgtype.Int4
	IsStored    pgtype.Bool
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (int32, error) {
	var itemId int32
	err := q.db.QueryRow(ctx, createItem,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.GroupID,
		arg.LocationID,
		arg.IsStored,
	).Scan(&itemId)
	if err != nil {
		return 0, err
	}
	return itemId, nil
}
