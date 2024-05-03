// Package locationDomain domain/locationDomain/domainLocation.go
package locationDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Location struct {
	LocationID int32
	TubID      pgtype.Int4
	ShelfID    pgtype.Int4
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type Repository interface {
	GetAllLocations() ([]*Location, error)
	GetLocationByID(id int32) (*Location, error)
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

func (s *Service) GetAllLocations() ([]*Location, error) {
	return s.repo.GetAllLocations()
}

func (s *Service) GetLocationByID(id int32) (*Location, error) {
	return s.repo.GetLocationByID(id)
}

func (r *Repo) GetLocationByID(id int32) (*Location, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetItem method to retrieve the Location from the database
	dbItem, err := q.GetLocation(context.Background(), id)
	if err != nil {
		return nil, err
	}

	// Map the db.Location to domain.Location
	location := &Location{
		LocationID: dbItem.LocationID,
		TubID:      dbItem.TubID,
		ShelfID:    dbItem.ShelfID,
		CreatedAt:  dbItem.CreatedAt,
		UpdatedAt:  dbItem.UpdatedAt,
	}
	return location, nil
}

func (r *Repo) GetAllLocations() ([]*Location, error) {
	// Create a new instance of the Queries struct
	q := db.New(r.db)

	// Call the GetAllLocations method to retrieve the Location from the database
	dbItems, err := q.GetAllLocations(context.Background())
	if err != nil {
		return nil, err
	}

	// Map the db.Location to domain.Location
	var locations []*Location

	for _, dbLocation := range dbItems {
		location := &Location{
			LocationID: dbLocation.LocationID,
			TubID:      dbLocation.TubID,
			ShelfID:    dbLocation.ShelfID,
			CreatedAt:  dbLocation.CreatedAt,
			UpdatedAt:  dbLocation.UpdatedAt,
		}
		locations = append(locations, location)
	}
	return locations, nil
}
