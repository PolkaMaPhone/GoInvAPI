package categoryDomain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetCategoryByID(id int32) (*Category, error) {
	args := m.Called(id)
	return args.Get(0).(*Category), args.Error(1)
}

func (m *MockRepository) GetAllCategories() ([]*Category, error) {
	args := m.Called()
	return args.Get(0).([]*Category), args.Error(1)
}

func TestService_GetCategoryByID(t *testing.T) {
	mockRepo := new(MockRepository)
	category := &Category{CategoryID: 1}
	mockRepo.On("GetCategoryByID", int32(1)).Return(category, nil)

	service := NewService(mockRepo)
	result, err := service.GetCategoryByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.CategoryID, result.CategoryID)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllCategories(t *testing.T) {
	mockRepo := new(MockRepository)
	categories := []*Category{{CategoryID: 1}, {CategoryID: 2}}
	mockRepo.On("GetAllCategories").Return(categories, nil)

	service := NewService(mockRepo)
	result, err := service.GetAllCategories()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(categories), len(result))

	mockRepo.AssertExpectations(t)
}
