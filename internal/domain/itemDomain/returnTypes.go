package itemDomain

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type PartialItem struct {
	ItemID    int32              `json:"ItemID"`
	UpdatedAt pgtype.Timestamptz `json:"UpdatedAt"`
	CreatedAt pgtype.Timestamptz `json:"CreatedAt"`
}

type DeleteResponse struct {
	ID      int32  `json:"id"`
	Message string `json:"message"`
}
