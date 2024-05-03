// Package groupDomain domain/groupDomain/domainGroup.go
package groupDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Group struct {
	GroupID     int32
	Name        string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type Repository interface {
	GetAllGroups() ([]*Group, error)
	GetGroupByID(id int32) (*Group, error)
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

func (s *Service) GetAllGroups() ([]*Group, error) {
	return s.repo.GetAllGroups()
}

func (s *Service) GetGroupByID(id int32) (*Group, error) {
	return s.repo.GetGroupByID(id)
}

func (r *Repo) GetGroupByID(id int32) (*Group, error) {
	q := db.New(r.db)
	dbItem, err := q.GetGroup(context.Background(), id)
	if err != nil {
		return nil, err
	}

	group := &Group{
		GroupID:     dbItem.GroupID,
		Name:        dbItem.Name,
		Description: dbItem.Description,
		CreatedAt:   dbItem.CreatedAt,
		UpdatedAt:   dbItem.UpdatedAt,
	}
	return group, nil
}

func (r *Repo) GetAllGroups() ([]*Group, error) {
	q := db.New(r.db)
	dbItem, err := q.GetAllGroups(context.Background())
	if err != nil {
		return nil, err
	}

	var groups []*Group
	for _, dbGroup := range dbItem {
		group := &Group{
			GroupID:     dbGroup.GroupID,
			Name:        dbGroup.Name,
			Description: dbGroup.Description,
			CreatedAt:   dbGroup.CreatedAt,
			UpdatedAt:   dbGroup.UpdatedAt,
		}
		groups = append(groups, group)
	}
	return groups, nil
}
