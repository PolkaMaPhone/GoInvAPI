// Package itemDomain domain/item/mappers.go
package itemDomain

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapDBItemToDomainItem(dbItem *db.Item) (*Item, error) {
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.Item to domain.Item: dbItem is nil")
		return nil, errors.New("dbItem cannot be nil")
	}
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
	}, nil
}

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

func convertToDTOItemWithGroup(domainItem *ItemWithGroup) *dto.ItemWithGroup {
	return &dto.ItemWithGroup{
		ItemID:           domainItem.ItemID,
		Name:             domainItem.Name,
		Description:      domainItem.Description,
		CategoryID:       domainItem.CategoryID,
		GroupID:          domainItem.GroupID,
		LocationID:       domainItem.LocationID,
		IsStored:         domainItem.IsStored,
		CreatedAt:        domainItem.CreatedAt,
		UpdatedAt:        domainItem.UpdatedAt,
		GroupName:        domainItem.GroupName,
		GroupDescription: domainItem.GroupDescription,
	}
}

// MapDBItemWithCategoryToDTO maps a db.GetItemWithCategoryRow to a dto.ItemWithCategory
func MapDBItemWithCategoryToDTO(dbItem *db.GetItemWithCategoryRow) *dto.ItemWithCategory {

	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetItemWithCategoryRow to dto.ItemWithCategory: dbItem is nil")
		return nil
	}
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
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetAllItemsWithCategoriesRow to dto.ItemWithCategory: dbItem is nil")
		return nil
	}
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

func MapDBAllItemsWithGroupsToDTO(dbItem *db.GetAllItemsWithGroupsRow) *dto.ItemWithGroup {
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetAllItemsWithGroupsRow to dto.ItemWithGroup: dbItem is nil")
		return nil
	}
	domainItem := &ItemWithGroup{
		ItemID:           dbItem.ItemID,
		Name:             dbItem.Name,
		Description:      dbItem.Description,
		CategoryID:       dbItem.CategoryID,
		GroupID:          dbItem.GroupID,
		LocationID:       dbItem.LocationID,
		IsStored:         dbItem.IsStored,
		CreatedAt:        dbItem.CreatedAt,
		UpdatedAt:        dbItem.UpdatedAt,
		GroupName:        dbItem.GroupName,
		GroupDescription: dbItem.GroupDescription,
	}
	if dbItem.GroupID.Valid == false {
		domainItem.GroupName = pgtype.Text{String: "Ungrouped", Valid: true}
		domainItem.GroupDescription = pgtype.Text{String: "An Ungrouped Item", Valid: true}
	}

	return convertToDTOItemWithGroup(domainItem)
}
