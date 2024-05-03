package groupDomain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetGroupByID(id int32) (*Group, error) {
	args := m.Called(id)
	return args.Get(0).(*Group), args.Error(1)
}

func (m *MockRepository) GetAllGroups() ([]*Group, error) {
	args := m.Called()
	return args.Get(0).([]*Group), args.Error(1)
}

func TestService_GetGroupByID(t *testing.T) {
	mockRepo := new(MockRepository)
	group := &Group{GroupID: 1}
	mockRepo.On("GetGroupByID", int32(1)).Return(group, nil)

	service := NewService(mockRepo)
	result, err := service.GetGroupByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, group.GroupID, result.GroupID)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllGroups(t *testing.T) {
	mockRepo := new(MockRepository)
	groups := []*Group{{GroupID: 1}, {GroupID: 2}}
	mockRepo.On("GetAllGroups").Return(groups, nil)

	service := NewService(mockRepo)
	result, err := service.GetAllGroups()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(groups), len(result))

	mockRepo.AssertExpectations(t)
}
