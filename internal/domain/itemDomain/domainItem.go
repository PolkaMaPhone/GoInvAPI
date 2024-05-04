// Package itemDomain domain/item/domainItem.go
package itemDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO - Implement GetAllItemsWithLocations
// TODO - Implement GetItemWithLocationByID

// TODO - Implement GetAllItemsWithDetails
// TODO - Implement GetItemWithDetailByID

// TODO - Maybe refactor by type into other files but the same package?

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

func (s *Service) GetItemByID(id int32) (*Item, error) {
	return s.repo.GetItemByID(id)
}

func (s *Service) GetAllItems() ([]*Item, error) {
	return s.repo.GetAllItems()
}

func (s *Service) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	return s.repo.GetItemByIDWithCategory(id)
}

func (s *Service) GetItemByIDWithGroup(id int32) (*dto.ItemWithGroup, error) {
	return s.repo.GetItemByIDWithGroup(id)
}

func (s *Service) GetItemByIDWithGroupAndCategory(id int32) (*dto.ItemWithGroupAndCategory, error) {
	return s.repo.GetItemByIDWithGroupAndCategory(id)
}

func (s *Service) GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error) {
	return s.repo.GetAllItemsWithCategories()
}

func (s *Service) GetAllItemsWithGroups() ([]*dto.ItemWithGroup, error) {
	return s.repo.GetAllItemsWithGroups()
}

func (s *Service) GetAllItemsWithGroupsAndCategories() ([]*dto.ItemWithGroupAndCategory, error) {
	return s.repo.GetAllItemsWithGroupsAndCategories()
}

func (r *Repo) GetItemByID(id int32) (*Item, error) {
	q := db.New(r.db)

	dbItem, err := q.GetItem(context.Background(), id)
	if err != nil {
		return nil, err
	}

	item, _ := MapDBItemToDomainItem(&dbItem)
	return item, nil
}

func (r *Repo) GetAllItems() ([]*Item, error) {
	q := db.New(r.db)

	dbItems, err := q.GetAllItems(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.Item to domain.Item
	var items []*Item
	for _, dbItem := range dbItems {
		item, _ := MapDBItemToDomainItem(&dbItem)
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
	item := MapDBItemWithCategoryToDTO(&dbItem)

	return item, nil
}

func (r *Repo) GetItemByIDWithGroup(id int32) (*dto.ItemWithGroup, error) {
	q := db.New(r.db)
	dbItem, err := q.GetItemWithGroup(context.Background(), id)
	if err != nil {
		return nil, err
	}
	item := MapDBItemWithGroupToDTO(&dbItem)
	return item, nil
}

func (r *Repo) GetItemByIDWithGroupAndCategory(id int32) (*dto.ItemWithGroupAndCategory, error) {
	q := db.New(r.db)
	dbItem, err := q.GetItemWithGroupAndCategory(context.Background(), id)
	if err != nil {
		return nil, err
	}
	item := MapDBItemWithGroupAndCategoryToDTO(&dbItem)
	return item, nil
}

//func (r *Repo) GetItemByIDWithLocation(id int32) (*dto.ItemWithLocation, error) {
//	q := db.New(r.db)
//	dbItem, err := q.GetItemWithLocation(context.Background(), id)
//	if err != nil {
//		return nil, err
//	}
//	item := MapDBItemWithLocationToDTO(&dbItem)
//	return item, nil
//}

func (r *Repo) GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllItemsWithCategories method to retrieve all items from the database
	dbItems, err := q.GetAllItemsWithCategories(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.ItemWithCategory to dto.ItemWithCategory
	var items []*dto.ItemWithCategory
	for _, dbItem := range dbItems {
		item := MapDBAllItemsWithCategoriesToDTO(&dbItem)
		items = append(items, item)
	}

	return items, nil
}

func (r *Repo) GetAllItemsWithGroups() ([]*dto.ItemWithGroup, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllItemsWithGroups method to retrieve all items from the database
	dbItems, err := q.GetAllItemsWithGroups(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.ItemWithGroup to dto.ItemWithGroup
	var items []*dto.ItemWithGroup
	for _, dbItem := range dbItems {
		item := MapDBAllItemsWithGroupsToDTO(&dbItem)
		items = append(items, item)
	}

	return items, nil
}

func (r *Repo) GetAllItemsWithGroupsAndCategories() ([]*dto.ItemWithGroupAndCategory, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllItemsWithGroupsAndCategories method to retrieve all items from the database
	dbItems, err := q.GetAllItemsWithGroupsAndCategories(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.ItemWithGroupAndCategory to dto.ItemWithGroupAndCategory
	var items []*dto.ItemWithGroupAndCategory
	for _, dbItem := range dbItems {
		item := MapDBAllItemsWithGroupsAndCategoriesToDTO(&dbItem)
		items = append(items, item)
	}

	return items, nil
}

//func (r *Repo) GetItemByIDWithLocation(id int32) (*dto.ItemWithLocation, error) {
//	q := db.New(r.db)
//	dbItem, err := q.GetItemWithLocation(context.Background(), id)
//	if err != nil {
//		return nil, err
//	}
//	item := MapDBItemWithLocationToDTO(&dbItem)
//	return item, nil
//}
