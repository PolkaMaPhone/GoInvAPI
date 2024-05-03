package locationDomain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetLocationByID(id int32) (*Location, error) {
	args := m.Called(id)
	return args.Get(0).(*Location), args.Error(1)
}

func (m *MockRepository) GetAllLocations() ([]*Location, error) {
	args := m.Called()
	return args.Get(0).([]*Location), args.Error(1)
}

func TestService_GetLocationByID(t *testing.T) {
	mockRepo := new(MockRepository)
	location := &Location{LocationID: 1}
	mockRepo.On("GetLocationByID", int32(1)).Return(location, nil)

	service := NewService(mockRepo)
	result, err := service.GetLocationByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, location.LocationID, result.LocationID)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllLocations(t *testing.T) {
	mockRepo := new(MockRepository)
	locations := []*Location{{LocationID: 1}, {LocationID: 2}}
	mockRepo.On("GetAllLocations").Return(locations, nil)

	service := NewService(mockRepo)
	result, err := service.GetAllLocations()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(locations), len(result))

	mockRepo.AssertExpectations(t)
}
