// Package category domain/category/category.go
package category

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category struct {
	CategoryID  int32
	Name        string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type Repository interface {
	GetAllCategories() ([]*Category, error)
	GetCategoryByID(id int32) (*Category, error)
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

func (s *Service) GetAllCategories() ([]*Category, error) {
	return s.repo.GetAllCategories()
}

func (s *Service) GetCategoryByID(id int32) (*Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (r *Repo) GetCategoryByID(id int32) (*Category, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetItem method to retrieve the Category from the database
	dbItem, err := q.GetCategory(context.Background(), id)
	if err != nil {
		return nil, err
	}

	// Map the db.Category to domain.Category
	category := &Category{
		CategoryID:  dbItem.CategoryID,
		Name:        dbItem.Name,
		Description: dbItem.Description,
		CreatedAt:   dbItem.CreatedAt,
		UpdatedAt:   dbItem.UpdatedAt,
	}
	return category, nil
}

func (r *Repo) GetAllCategories() ([]*Category, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllCategories method to retrieve the Category from the database
	dbItem, err := q.GetAllCategories(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.Category to domain.Category
	var categories []*Category

	for _, dbCategory := range dbItem {
		category := &Category{
			CategoryID:  dbCategory.CategoryID,
			Name:        dbCategory.Name,
			Description: dbCategory.Description,
			CreatedAt:   dbCategory.CreatedAt,
			UpdatedAt:   dbCategory.UpdatedAt,
		}
		categories = append(categories, category)
	}
	return categories, nil
}
