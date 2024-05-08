package itemInterface

import (
	"context"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/internal/application/dto"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/dbconn"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http/httptest"
	"strconv"
)

type MockService struct {
	mock.Mock
}

func getExpectedErrorMessage(err error) string {
	switch e := err.(type) {
	case *utils.InvalidRouteError:
		return fmt.Sprintf("invalid route '%s'", e.Route)
	case *utils.MethodNotAllowedError:
		return fmt.Sprintf("method '%s' is not allowed for route '%s'", e.Method, e.Route)
	case *utils.NoResultsForParameterError:
		return fmt.Sprintf("the parameter '%s' with id '%s' returned no results", e.ParameterName, e.ID)
	case *utils.InvalidParameterError:
		return fmt.Sprintf("invalid parameter '%s'", e.ParameterName)
	case *utils.ServerErrorType:
		return "internal server error"
	case *strconv.NumError:
		return fmt.Sprintf("invalid parameter: cannot parse '%s' as %s", e.Num, e.Func)
	default:
		return fmt.Sprintf("unexpected error type %T: %v", e, e)
	}
}

func initializeItemTestServer() *httptest.Server {
	config, err := dbconn.LoadConfigFile()
	if err != nil {
		log.Fatalf("Unable to load config: %v\n", err)
	}

	dbtest := &dbconn.PgxDB{}
	_, err = dbconn.GetPoolInstance(config, dbtest)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	itemRepo := itemDomain.NewRepository(dbtest.Pool)
	itemService := itemDomain.NewService(itemRepo)
	itemHandler := NewItemHandler(itemService)

	router := chi.NewRouter()
	r := &customRouter.CustomRouter{
		Mux: router,
	}
	itemHandler.HandleRoutes(r)

	server := httptest.NewServer(router)
	return server
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
