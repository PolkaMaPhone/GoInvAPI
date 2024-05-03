package itemDomain

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetItemByID(id int32) (*Item, error) {
	args := m.Called(id)
	return args.Get(0).(*Item), args.Error(1)
}

func (m *MockRepository) GetAllItems() ([]*Item, error) {
	args := m.Called()
	return args.Get(0).([]*Item), args.Error(1)
}
func (m *MockRepository) GetItemByIDWithCategory(id int32) (*dto.ItemWithCategory, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ItemWithCategory), args.Error(1)
}

func (m *MockRepository) GetAllItemsWithCategories() ([]*dto.ItemWithCategory, error) {
	args := m.Called()
	return args.Get(0).([]*dto.ItemWithCategory), args.Error(1)
}

func TestService_GetItemByID(t *testing.T) {
	mockRepo := new(MockRepository)
	item := &Item{ItemID: 1}
	mockRepo.On("GetItemByID", int32(1)).Return(item, nil)

	service := NewService(mockRepo)
	result, err := service.GetItemByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, item.ItemID, result.ItemID)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllItems(t *testing.T) {
	mockRepo := new(MockRepository)
	items := []*Item{{ItemID: 1}, {ItemID: 2}}
	mockRepo.On("GetAllItems").Return(items, nil)

	service := NewService(mockRepo)
	result, err := service.GetAllItems()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(items), len(result))

	mockRepo.AssertExpectations(t)
}

func TestService_GetItemByIDWithCategory(t *testing.T) {
	mockRepo := new(MockRepository)
	item := &dto.ItemWithCategory{
		ItemID:              1,
		Name:                "Item1",
		Description:         pgtype.Text{String: "Description1", Valid: true},
		CategoryID:          pgtype.Int4{Int32: 1, Valid: true},
		GroupID:             pgtype.Int4{Int32: 1, Valid: true},
		LocationID:          pgtype.Int4{Int32: 1, Valid: true},
		IsStored:            pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:           pgtype.Timestamptz{},
		UpdatedAt:           pgtype.Timestamptz{},
		CategoryName:        pgtype.Text{String: "Category1", Valid: true},
		CategoryDescription: pgtype.Text{String: "CategoryDescription1", Valid: true},
	}
	mockRepo.On("GetItemByIDWithCategory", int32(1)).Return(item, nil)

	service := NewService(mockRepo)
	result, err := service.GetItemByIDWithCategory(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, item.ItemID, result.ItemID)
	assert.Equal(t, item.Name, result.Name)
	assert.Equal(t, item.Description.String, result.Description.String)
	assert.Equal(t, item.CategoryID.Int32, result.CategoryID.Int32)
	assert.Equal(t, item.GroupID.Int32, result.GroupID.Int32)
	assert.Equal(t, item.LocationID.Int32, result.LocationID.Int32)
	assert.Equal(t, item.IsStored.Bool, result.IsStored.Bool)
	assert.Equal(t, item.CategoryName, result.CategoryName)
	assert.Equal(t, item.CategoryDescription.String, result.CategoryDescription.String)

	mockRepo.AssertExpectations(t)

}

func TestService_GetAllItemsWithCategory(t *testing.T) {
	mockRepo := new(MockRepository)
	items := []*dto.ItemWithCategory{
		{
			ItemID:              1,
			Name:                "Item1",
			Description:         pgtype.Text{String: "Description1", Valid: true},
			CategoryID:          pgtype.Int4{Int32: 1, Valid: true},
			GroupID:             pgtype.Int4{Int32: 1, Valid: true},
			LocationID:          pgtype.Int4{Int32: 1, Valid: true},
			IsStored:            pgtype.Bool{Bool: true, Valid: true},
			CreatedAt:           pgtype.Timestamptz{},
			UpdatedAt:           pgtype.Timestamptz{},
			CategoryName:        pgtype.Text{String: "Category1", Valid: true},
			CategoryDescription: pgtype.Text{String: "CategoryDescription1", Valid: true},
		},
		{
			ItemID:              2,
			Name:                "Item2",
			Description:         pgtype.Text{String: "Description2", Valid: true},
			CategoryID:          pgtype.Int4{Int32: 2, Valid: true},
			GroupID:             pgtype.Int4{Int32: 2, Valid: true},
			LocationID:          pgtype.Int4{Int32: 2, Valid: true},
			IsStored:            pgtype.Bool{Bool: true, Valid: true},
			CreatedAt:           pgtype.Timestamptz{},
			UpdatedAt:           pgtype.Timestamptz{},
			CategoryName:        pgtype.Text{String: "Category2", Valid: true},
			CategoryDescription: pgtype.Text{String: "CategoryDescription2", Valid: true},
		},
	}
	mockRepo.On("GetAllItemsWithCategories").Return(items, nil)

	service := NewService(mockRepo)
	result, err := service.GetAllItemsWithCategories()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(items), len(result))

	for i, item := range result {
		assert.Equal(t, items[i].ItemID, item.ItemID)
		assert.Equal(t, items[i].Name, item.Name)
		assert.Equal(t, items[i].Description.String, item.Description.String)
		assert.Equal(t, items[i].CategoryID.Int32, item.CategoryID.Int32)
		assert.Equal(t, items[i].GroupID.Int32, item.GroupID.Int32)
		assert.Equal(t, items[i].LocationID.Int32, item.LocationID.Int32)
		assert.Equal(t, items[i].IsStored.Bool, item.IsStored.Bool)
		assert.Equal(t, items[i].CategoryName, item.CategoryName)
		assert.Equal(t, items[i].CategoryDescription.String, item.CategoryDescription.String)
	}

	mockRepo.AssertExpectations(t)
}
