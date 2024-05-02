package item

import (
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
