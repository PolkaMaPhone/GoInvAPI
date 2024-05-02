// Package itemDomain domain/item/domainItem.go
package itemDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Item struct {
	ItemID      int32
	Name        string
	Description pgtype.Text
	CategoryID  pgtype.Int4
	GroupID     pgtype.Int4
	LocationID  pgtype.Int4
	IsStored    pgtype.Bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type ItemWithCategory struct {
	ItemID              int32
	Name                string
	Description         pgtype.Text
	CategoryID          pgtype.Int4
	GroupID             pgtype.Int4
	LocationID          pgtype.Int4
	IsStored            pgtype.Bool
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	CategoryName        string
	CategoryDescription pgtype.Text
}

type Repository interface {
	GetAllItems() ([]*Item, error)
	GetItemByID(id int32) (*Item, error)
	GetAllItemsWithCategory() ([]*dto.ItemWithCategory, error)
	GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error)
}

type Service struct {
	repo Repository
}

type Repo struct {
	db *pgxpool.Pool
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &Repo{db: db}
}

func (s *Service) GetAllItems() ([]*Item, error) {
	return s.repo.GetAllItems()
}

func (s *Service) GetItemByID(id int32) (*Item, error) {
	return s.repo.GetItemByID(id)
}

func (s *Service) GetAllItemsWithCategory() ([]*dto.ItemWithCategory, error) {
	return s.repo.GetAllItemsWithCategory()
}

func (s *Service) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	return s.repo.GetItemByIDWithCategory(id)
}

func (r *Repo) GetItemByID(id int32) (*Item, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetItem method to retrieve the item from the database
	dbItem, err := q.GetItem(context.Background(), id)
	if err != nil {
		return nil, err
	}

	// Map the db.Item to domain.Item
	item := &Item{
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
	return item, nil
}

func (r *Repo) GetAllItems() ([]*Item, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllItems method to retrieve all items from the database
	dbItems, err := q.GetAllItems(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.Item to domain.Item
	var items []*Item
	for _, dbItem := range dbItems {
		item := &Item{
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
		items = append(items, item)
	}

	return items, nil
}

func (r *Repo) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetItem method to retrieve the item from the database
	dbItem, err := q.GetItemWithCategory(context.Background(), id)
	if err != nil {
		return nil, err
	}

	// Map the db.Item to domain.Item
	item := &dto.ItemWithCategory{
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
	//if the CategoryID is null, set CategoryName to "Uncategorized"
	if dbItem.CategoryID.Valid == false {
		item.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
		item.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
	}

	return item, nil
}

func (r *Repo) GetAllItemsWithCategory() ([]*dto.ItemWithCategory, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllItemsWithCategory method to retrieve all items from the database
	dbItems, err := q.GetAllItemsWithCategories(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.ItemWithCategory to dto.ItemWithCategory
	var items []*dto.ItemWithCategory
	for _, dbItem := range dbItems {
		item := &dto.ItemWithCategory{
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
		//if the CategoryID is null, set CategoryName to "Uncategorized"
		if dbItem.CategoryID.Valid == false {
			item.CategoryName = pgtype.Text{String: "Uncategorized", Valid: true}
			item.CategoryDescription = pgtype.Text{String: "An Uncategorized Item", Valid: true}
		}

		items = append(items, item)
	}

	return items, nil
}
