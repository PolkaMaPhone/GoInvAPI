package itemDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &Repo{db: db}
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

func (r *Repo) CreateItem(ctx context.Context, arg db.CreateItemParams) (*PartialItem, error) {
	q := db.New(r.db)
	logging.InfoLogger.Printf("Creating item %v", arg)

	itemID, err := q.CreateItem(ctx, arg)
	if err != nil {
		return nil, err
	}

	dbItem, err := q.GetItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	partialItem, _ := MapDBItemToPartialItem(&dbItem)
	return partialItem, nil
}

func (r *Repo) UpdateItem(ctx context.Context, arg db.UpdateItemParams) (*PartialItem, error) {
	q := db.New(r.db)

	_, err := q.UpdateItem(ctx, arg)
	if err != nil {
		return nil, err
	}

	dbItem, err := q.GetItem(ctx, arg.ItemID)
	if err != nil {
		return nil, err
	}

	returnItem, _ := MapDBItemToPartialItem(&dbItem)
	return returnItem, nil
}

func (r *Repo) DeleteItem(id int32) error {
	q := db.New(r.db)

	_, err := q.DeleteItem(context.Background(), id)
	return err
}
