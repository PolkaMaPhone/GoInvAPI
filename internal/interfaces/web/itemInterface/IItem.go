package itemInterface

import (
	"errors"
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
)

// TODO - Handle Item with Location

type Handler struct {
	service *itemDomain.Service
}

func NewItemHandler(s *itemDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.LoggingMiddleware("INFO"))
	apiRouter.HandleFunc("/items/{item_id}", h.HandleGet).Methods("GET")
	apiRouter.HandleFunc("/items/{item_id}/with_category", h.HandleGetWithCategory).Methods("GET")
	apiRouter.HandleFunc("/items/{item_id}/with_group", h.HandleGetWithGroup).Methods("GET")
	apiRouter.HandleFunc("/items/{item_id}/with_group_and_category", h.HandleGetWithGroupAndCategory).Methods("GET")

	apiRouter.HandleFunc("/items", h.HandleGetAll).Methods("GET")
	apiRouter.HandleFunc("/items_with_category", h.HandleGetAllWithCategories).Methods("GET")
	apiRouter.HandleFunc("/items_with_group", h.HandleGetAllWithGroups).Methods("GET")
	apiRouter.HandleFunc("/items_with_group_and_category", h.HandleGetAllWithGroupsAndCategories).Methods("GET")
}

func getItemIDFromRequest(w http.ResponseWriter, r *http.Request) (int32, error) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		utils.HandleHTTPError(w, &utils.InvalidParameterError{ParameterName: "item_id"}, http.StatusBadRequest)
		return 0, err
	}
	return int32(itemID), nil
}

func (h *Handler) handleGetItemErrors(w http.ResponseWriter, err error, item interface{}, itemID int32) bool {
	if err != nil {
		// Check if the error is a pgx.ErrNoRows error
		if errors.Is(err, pgx.ErrNoRows) {
			httpError := &utils.NoResultsForParameterError{ParameterName: "item_id", ID: strconv.Itoa(int(itemID)), StatusCode: http.StatusNoContent}
			utils.HandleHTTPError(w, httpError, httpError.StatusCode)
		} else {
			// For all other errors, return a 500 status code and a generic server error message
			utils.HandleHTTPError(w, &utils.ServerErrorType{}, http.StatusInternalServerError)
		}
		return true
	}

	if item == nil {
		httpError := &utils.NoResultsForParameterError{ParameterName: "item_id", ID: strconv.Itoa(int(itemID)), StatusCode: http.StatusNoContent}
		utils.HandleHTTPError(w, httpError, httpError.StatusCode)
		return true
	}

	return false
}

func (h *Handler) handleGetAllErrors(w http.ResponseWriter, err error, items interface{}) bool {
	if err != nil {
		// Check if the error is a sql.ErrNoRows error
		if errors.Is(err, pgx.ErrNoRows) {
			httpError := &utils.NoResultsForParameterError{ParameterName: "items", ID: "all", StatusCode: http.StatusNoContent}
			utils.HandleHTTPError(w, httpError, httpError.StatusCode)
		} else {
			// For all other errors, return a 500 status code and a generic server error message
			utils.HandleHTTPError(w, &utils.ServerErrorType{}, http.StatusInternalServerError)
		}
		return true
	}

	// Check if items is a slice
	if itemsSlice, ok := items.([]interface{}); ok {
		if itemsSlice == nil || len(itemsSlice) == 0 {
			httpError := &utils.NoResultsForParameterError{ParameterName: "items", ID: "all", StatusCode: http.StatusNoContent}
			utils.HandleHTTPError(w, httpError, httpError.StatusCode)
			return true
		}
	}

	return false
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	itemID, err := getItemIDFromRequest(w, r)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByID(itemID)
	if h.handleGetItemErrors(w, err, foundItem, itemID) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetWithCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := getItemIDFromRequest(w, r)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithCategory(itemID)
	if h.handleGetItemErrors(w, err, foundItem, itemID) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetWithGroup(w http.ResponseWriter, r *http.Request) {
	itemID, err := getItemIDFromRequest(w, r)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroup(itemID)
	if h.handleGetItemErrors(w, err, foundItem, itemID) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetWithGroupAndCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := getItemIDFromRequest(w, r)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroupAndCategory(itemID)
	if h.handleGetItemErrors(w, err, foundItem, itemID) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItems()
	if h.handleGetAllErrors(w, err, items) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithCategories()
	if h.handleGetAllErrors(w, err, items) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithGroups(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroups()
	if h.handleGetAllErrors(w, err, items) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithGroupsAndCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroupsAndCategories()
	if h.handleGetAllErrors(w, err, items) {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}
