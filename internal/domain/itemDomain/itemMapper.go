// Package itemDomain domain/item/mappers.go
package itemDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// MapDBItemToDomainItem maps a db.Item to a domain.Item
func MapDBItemToDomainItem(dbItem *db.Item) *Item {
	return &Item{
		ItemID:      dbItem.ItemID,
		Name:        dbItem.Name,
		Description: dbItem.Description,
		CategoryID:  dbItem.CategoryID,
		GroupID:     dbItem.GroupID,
		LocationID:  dbItem.LocationID,
		IsStored:    dbItem.IsStored,
		CreatedAt:   dbItem.CreatedAt,
		UpdatedAt:   dbItem.UpdatedAt,
	}
}

// explicit conversion function for safety and clarity
func convertToDTOItemWithCategory(domainItem *ItemWithCategory) *dto.ItemWithCategory {
	return &dto.ItemWithCategory{
		ItemID:              domainItem.ItemID,
		Name:                domainItem.Name,
		Description:         domainItem.Description,
		CategoryID:          domainItem.CategoryID,
		GroupID:             domainItem.GroupID,
		LocationID:          domainItem.LocationID,
		IsStored:            domainItem.IsStored,
		CreatedAt:           domainItem.CreatedAt,
		UpdatedAt:           domainItem.UpdatedAt,
		CategoryName:        domainItem.CategoryName,
		CategoryDescription: domainItem.CategoryDescription,
	}
}

// MapDBItemWithCategoryToDTO maps a db.GetItemWithCategoryRow to a dto.ItemWithCategory
func MapDBItemWithCategoryToDTO(dbItem *db.GetItemWithCategoryRow) *dto.ItemWithCategory {
	domainItem := &ItemWithCategory{
		ItemID:              dbItem.ItemID,
		Name:                dbItem.Name,
		Description:         dbItem.Description,
		CategoryID:          dbItem.CategoryID,
		GroupID:             dbItem.GroupID,
		LocationID:          dbItem.LocationID,
		IsStored:            dbItem.IsStored,
		CreatedAt:           dbItem.CreatedAt,
		UpdatedAt:           dbItem.UpdatedAt,
		CategoryName:        dbItem.CategoryName,
		CategoryDescription: dbItem.CategoryDescription,
	}
	if dbItem.CategoryID.Valid == false {
		domainItem.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
		domainItem.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
	}

	return convertToDTOItemWithCategory(domainItem)
}

// MapDBAllItemsWithCategoriesToDTO maps a db.GetAllItemsWithCategoriesRow to a dto.ItemWithCategory
func MapDBAllItemsWithCategoriesToDTO(dbItem *db.GetAllItemsWithCategoriesRow) *dto.ItemWithCategory {
	domainItem := &ItemWithCategory{
		ItemID:              dbItem.ItemID,
		Name:                dbItem.Name,
		Description:         dbItem.Description,
		CategoryID:          dbItem.CategoryID,
		GroupID:             dbItem.GroupID,
		LocationID:          dbItem.LocationID,
		IsStored:            dbItem.IsStored,
		CreatedAt:           dbItem.CreatedAt,
		UpdatedAt:           dbItem.UpdatedAt,
		CategoryName:        dbItem.CategoryName,
		CategoryDescription: dbItem.CategoryDescription,
	}
	if dbItem.CategoryID.Valid == false {
		domainItem.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
		domainItem.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
	}

	return convertToDTOItemWithCategory(domainItem)
}
