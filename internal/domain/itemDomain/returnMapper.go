package itemDomain

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
)

func MapDBItemToPartialItem(dbItem *db.Item) (*PartialItem, error) {
	if dbItem == nil {
		logging.ErrorLogger.Println("Failed to map db.Item to domain.Item: dbItem is nil")
		return nil, errors.New("dbItem cannot be nil")
	}
	return &PartialItem{
		ItemID:    dbItem.ItemID,
		CreatedAt: dbItem.CreatedAt,
		UpdatedAt: dbItem.UpdatedAt,
	}, nil
}
