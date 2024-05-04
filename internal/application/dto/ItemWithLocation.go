package dto

import "github.com/jackc/pgx/v5/pgtype"

type ItemWithLocation struct {
	ItemID      int32
	Name        string
	Description pgtype.Text
	CategoryID  pgtype.Int4
	GroupID     pgtype.Int4
	LocationID  pgtype.Int4
	IsStored    pgtype.Bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	TubLabel    pgtype.Text
	ShelfLabel  pgtype.Text
}
