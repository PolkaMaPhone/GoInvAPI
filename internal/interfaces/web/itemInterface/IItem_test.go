package itemInterface

import (
	"context"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetItemByIDWithGroup(id int32) (*dto.ItemWithGroup, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithGroup), args.Error(1)
}

func (m *MockService) GetItemByIDWithGroupAndCategory(id int32) (*dto.ItemWithGroupAndCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithGroupAndCategory), args.Error(1)
}

func (m *MockService) GetAllItemsWithGroups() ([]*dto.ItemWithGroup, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithGroup), args.Error(1)
}

func (m *MockService) GetAllItemsWithGroupsAndCategories() ([]*dto.ItemWithGroupAndCategory, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithGroupAndCategory), args.Error(1)
}

func (m *MockService) GetItemByID(id int32) (*itemDomain.Item, error) {
	args := m.Called(id)
	item, ok := args.Get(0).(*itemDomain.Item)
	if !ok {
		return nil, args.Error(1)
	}
	return item, args.Error(1)
}

func (m *MockService) GetAllItems() ([]*itemDomain.Item, error) {
	args := m.Called()
	return args.Get(0).([]*itemDomain.Item), args.Error(1)
}

func (m *MockService) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithCategory), args.Error(1)
}

func (m *MockService) GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithCategory), args.Error(1)
}

func (m *MockService) CreateItem(ctx context.Context, arg db.CreateItemParams) (*itemDomain.PartialItem, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(*itemDomain.PartialItem), args.Error(1)
}

func (m *MockService) DeleteItem(id int32) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockService) UpdateItem(ctx context.Context, arg db.UpdateItemParams) (*itemDomain.PartialItem, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(*itemDomain.PartialItem), args.Error(1)
}
