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

func convertToDTOItemWithGroupAndCategory(domainItem *ItemWithGroupAndCategory) *dto.ItemWithGroupAndCategory {
	return &dto.ItemWithGroupAndCategory{
		ItemID:              domainItem.ItemID,
		Name:                domainItem.Name,
		Description:         domainItem.Description,
		CategoryID:          domainItem.CategoryID,
		GroupID:             domainItem.GroupID,
		LocationID:          domainItem.LocationID,
		IsStored:            domainItem.IsStored,
		CreatedAt:           domainItem.CreatedAt,
		UpdatedAt:           domainItem.UpdatedAt,
		GroupName:           domainItem.GroupName,
		GroupDescription:    domainItem.GroupDescription,
		CategoryName:        domainItem.CategoryName,
		CategoryDescription: domainItem.CategoryDescription,
	}
}

func convertToDTOItemWithLocation(domainItem *ItemWithLocation) *dto.ItemWithLocation {
	return &dto.ItemWithLocation{
		ItemID:      domainItem.ItemID,
		Name:        domainItem.Name,
		Description: domainItem.Description,
		CategoryID:  domainItem.CategoryID,
		GroupID:     domainItem.GroupID,
		LocationID:  domainItem.LocationID,
		IsStored:    domainItem.IsStored,
		CreatedAt:   domainItem.CreatedAt,
		UpdatedAt:   domainItem.UpdatedAt,
		ShelfLabel:  domainItem.ShelfLabel,
		TubLabel:    domainItem.TubLabel,
	}
}

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

func MapDBItemWithGroupToDTO(dbItem *db.GetItemWithGroupRow) *dto.ItemWithGroup {
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetItemWithGroupRow to dto.ItemWithGroup: dbItem is nil")
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

func MapDBItemWithGroupAndCategoryToDTO(dbItem *db.GetItemWithGroupAndCategoryRow) *dto.ItemWithGroupAndCategory {
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetItemWithGroupAndCategoryRow to dto.ItemWithGroupAndCategory: dbItem is nil")
		return nil
	}
	domainItem := &ItemWithGroupAndCategory{
		ItemID:              dbItem.ItemID,
		Name:                dbItem.Name,
		Description:         dbItem.Description,
		CategoryID:          dbItem.CategoryID,
		GroupID:             dbItem.GroupID,
		LocationID:          dbItem.LocationID,
		IsStored:            dbItem.IsStored,
		CreatedAt:           dbItem.CreatedAt,
		UpdatedAt:           dbItem.UpdatedAt,
		GroupName:           dbItem.GroupName,
		GroupDescription:    dbItem.GroupDescription,
		CategoryName:        dbItem.CategoryName,
		CategoryDescription: dbItem.CategoryDescription,
	}
	if dbItem.GroupID.Valid == false {
		domainItem.GroupName = pgtype.Text{String: "Ungrouped", Valid: true}
		domainItem.GroupDescription = pgtype.Text{String: "An Ungrouped Item", Valid: true}
	}
	if dbItem.CategoryID.Valid == false {
		domainItem.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
		domainItem.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
	}
	return convertToDTOItemWithGroupAndCategory(domainItem)
}

//func MapDBItemWithLocationToDTO(dbItem *db.GetItemWithLocationRow) *dto.ItemWithLocation {
//	if dbItem == nil {
//		middleware.ErrorLogger.Println("Failed to map db.GetItemWithLocationRow to dto.ItemWithLocation: dbItem is nil")
//		return nil
//	}
//	domainItem := &ItemWithLocation{
//		ItemID:      dbItem.ItemID,
//		Name:        dbItem.Name,
//		Description: dbItem.Description,
//		CategoryID:  dbItem.CategoryID,
//		GroupID:     dbItem.GroupID,
//		LocationID:  dbItem.LocationID,
//		IsStored:    dbItem.IsStored,
//		CreatedAt:   dbItem.CreatedAt,
//		UpdatedAt:   dbItem.UpdatedAt,
//		TubLabel:    dbItem.TubLabel,
//		ShelfLabel:  dbItem.ShelfLabel,
//	}
//	return convertToDTOItemWithLocation(domainItem)
//}

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

func MapDBAllItemsWithGroupsAndCategoriesToDTO(dbItem *db.GetAllItemsWithGroupsAndCategoriesRow) *dto.ItemWithGroupAndCategory {
	if dbItem == nil {
		middleware.ErrorLogger.Println("Failed to map db.GetAllItemsWithGroupsAndCategoriesRow to dto.ItemWithGroupAndCategory: dbItem is nil")
		return nil
	}
	domainItem := &ItemWithGroupAndCategory{
		ItemID:              dbItem.ItemID,
		Name:                dbItem.Name,
		Description:         dbItem.Description,
		CategoryID:          dbItem.CategoryID,
		GroupID:             dbItem.GroupID,
		LocationID:          dbItem.LocationID,
		IsStored:            dbItem.IsStored,
		CreatedAt:           dbItem.CreatedAt,
		UpdatedAt:           dbItem.UpdatedAt,
		GroupName:           dbItem.GroupName,
		GroupDescription:    dbItem.GroupDescription,
		CategoryName:        dbItem.CategoryName,
		CategoryDescription: dbItem.CategoryDescription,
	}
	if dbItem.GroupID.Valid == false {
		domainItem.GroupName = pgtype.Text{String: "Ungrouped", Valid: true}
		domainItem.GroupDescription = pgtype.Text{String: "An Ungrouped Item", Valid: true}
	}
	if dbItem.CategoryID.Valid == false {
		domainItem.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
		domainItem.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
	}

	return convertToDTOItemWithGroupAndCategory(domainItem)
}

//func MapDBAllItemsWithLocationsToDTO(dbItem *db.GetAllItemsWithLocationsRow) *dto.ItemWithLocation {
//	if dbItem == nil {
//		middleware.ErrorLogger.Println("Failed to map db.GetAllItemsWithLocationsRow to dto.ItemWithLocation: dbItem is nil")
//		return nil
//	}
//	domainItem := &ItemWithLocation{
//		ItemID:      dbItem.ItemID,
//		Name:        dbItem.Name,
//		Description: dbItem.Description,
//		CategoryID:  dbItem.CategoryID,
//		GroupID:     dbItem.GroupID,
//		LocationID:  dbItem.LocationID,
//		IsStored:    dbItem.IsStored,
//		CreatedAt:   dbItem.CreatedAt,
//		UpdatedAt:   dbItem.UpdatedAt,
//		TubLabel:    dbItem.TubLabel,
//		ShelfLabel:  dbItem.ShelfLabel,
//	}
//
//	return convertToDTOItemWithLocation(domainItem)
//}
