package dto

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ItemWithGroup struct {
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
