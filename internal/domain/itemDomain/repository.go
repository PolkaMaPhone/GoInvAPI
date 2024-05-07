package itemDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
)

type Repository interface {
	GetItemByID(id int32) (*Item, error)
	GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error)
	GetItemByIDWithGroup(id int32) (*dto.ItemWithGroup, error)
	GetItemByIDWithGroupAndCategory(id int32) (*dto.ItemWithGroupAndCategory, error)
	//GetItemByIDWithLocation(id int32) (*dto.ItemWithLocation, error)

	GetAllItems() ([]*Item, error)
	GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error)
	GetAllItemsWithGroups() ([]*dto.ItemWithGroup, error)
	GetAllItemsWithGroupsAndCategories() ([]*dto.ItemWithGroupAndCategory, error)
	//GetAllItemsWithLocations() ([]*dto.ItemWithLocation, error)

	CreateItem(context.Context, db.CreateItemParams) (*PartialItem, error)
	DeleteItem(id int32) error
	UpdateItem(context.Context, db.UpdateItemParams) (*PartialItem, error)
}
