package itemDomain

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
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

func (s *Service) CreateItem(ctx context.Context, arg db.CreateItemParams) (*PartialItem, error) {
	return s.repo.CreateItem(ctx, arg)
}

func (s *Service) DeleteItem(id int32) error {
	return s.repo.DeleteItem(id)
}

func (s *Service) UpdateItem(ctx context.Context, arg db.UpdateItemParams) (*PartialItem, error) {
	return s.repo.UpdateItem(ctx, arg)
}
