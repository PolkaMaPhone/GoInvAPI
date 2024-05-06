package locationInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/locationDomain"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetLocationByID(id int32) (*locationDomain.Location, error) {
	args := m.Called(id)
	location, ok := args.Get(0).(*locationDomain.Location)
	if !ok {
		return nil, args.Error(1)
	}
	return location, args.Error(1)
}

func (m *MockService) GetAllLocations() ([]*locationDomain.Location, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*locationDomain.Location), args.Error(1)
}
