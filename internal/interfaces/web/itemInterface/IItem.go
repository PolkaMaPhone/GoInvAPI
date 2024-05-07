package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validation"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// TODO - Handle Item with Location

const idParameterName = "item_id"

type Handler struct {
	service *itemDomain.Service
}

func NewItemHandler(s *itemDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(apiRouter *customRouter.CustomRouter) {
	apiRouter.Route("/api/items", func(r chi.Router) {
		r.Route("/{item_id}", func(r chi.Router) {
			r.Use(validation.ValidateMethod(http.MethodGet, http.MethodPut, http.MethodDelete))
			r.Get("/", h.HandleGet)
			r.Get("/with_category", h.HandleGetWithCategory)
			r.Get("/with_group", h.HandleGetWithGroup)
			r.Get("/with_group_and_category", h.HandleGetWithGroupAndCategory)
			r.Delete("/", h.HandleDelete)
			r.Put("/", h.HandlePut)
		})

		r.With(validation.ValidateMethod(http.MethodPost)).Post("/", h.HandlePost)

		r.With(validation.ValidateMethod(http.MethodGet)).Get("/", h.HandleGetAll)
		r.With(validation.ValidateMethod(http.MethodGet)).Get("/with_category", h.HandleGetAllWithCategories)
		r.With(validation.ValidateMethod(http.MethodGet)).Get("/with_group", h.HandleGetAllWithGroups)
		r.With(validation.ValidateMethod(http.MethodGet)).Get("/with_group_and_category", h.HandleGetAllWithGroupsAndCategories)
	})
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByID(itemID)
	if utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	item, err := utils.DecodeItemFromRequest(w, r)
	if err != nil {
		logging.ErrorLogger.Printf("Error decoding item from request: %v", err)
		return
	}

	createdItem, err := h.service.CreateItem(r.Context(), db.CreateItemParams{
		Name:        item.Name,
		Description: item.Description,
		CategoryID:  item.CategoryID,
		GroupID:     item.GroupID,
		LocationID:  item.LocationID,
		IsStored:    item.IsStored,
	})
	if err != nil {
		logging.ErrorLogger.Printf("Error creating item: %v", err)
		return
	}
	logging.InfoLogger.Printf("Item created: %v", createdItem)
	utils.RespondWithJSON(w, http.StatusCreated, createdItem)
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	err = h.service.DeleteItem(itemID)
	if err != nil {
		logging.ErrorLogger.Printf("Error deleting item: %v", err)
		return
	}
	logging.InfoLogger.Printf("Item deleted: %v", itemID)

	response := itemDomain.DeleteResponse{
		ID:      itemID,
		Message: "Item Successfully Deleted",
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *Handler) HandlePut(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	// Get the existing item
	existingItem, err := h.service.GetItemByID(itemID)
	if err != nil {
		logging.ErrorLogger.Printf("Error getting item: %v", err)
		return
	}

	newItem, err := utils.DecodeItemFromRequest(w, r)
	if err != nil {
		logging.ErrorLogger.Printf("Error decoding item from request: %v", err)
		return
	}
	logging.InfoLogger.Printf("New item: %v", newItem)

	if newItem.Name == "" {
		logging.InfoLogger.Printf("Name is empty, using existing name: %v", existingItem.Name)
		newItem.Name = existingItem.Name
	}
	if newItem.Description.String == "" {
		logging.InfoLogger.Printf("Description is empty, using existing description: %v", existingItem.Description)
		newItem.Description = existingItem.Description
	}
	if newItem.CategoryID.Valid == false {
		logging.InfoLogger.Printf("CategoryID is empty, using existing CategoryID: %v", existingItem.CategoryID)
		newItem.CategoryID = existingItem.CategoryID
	}
	if newItem.GroupID.Valid == false {
		logging.InfoLogger.Printf("GroupID is empty, using existing GroupID: %v", existingItem.GroupID)
		newItem.GroupID = existingItem.GroupID

	}
	if newItem.LocationID.Valid == false {
		logging.InfoLogger.Printf("LocationID is empty, using existing LocationID: %v", existingItem.LocationID)
		newItem.LocationID = existingItem.LocationID
	}
	if newItem.IsStored.Valid == false {
		logging.InfoLogger.Printf("IsStored is empty, using existing IsStored: %v", existingItem.IsStored)
		newItem.IsStored = existingItem.IsStored
	}

	// Update the item in the database with the new struct
	updatedItem, err := h.service.UpdateItem(r.Context(), db.UpdateItemParams{
		ItemID:      itemID,
		Name:        newItem.Name,
		Description: newItem.Description,
		CategoryID:  newItem.CategoryID,
		GroupID:     newItem.GroupID,
		LocationID:  newItem.LocationID,
		IsStored:    newItem.IsStored,
	})
	if err != nil {
		logging.ErrorLogger.Printf("Error updating item: %v", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedItem)
}

func (h *Handler) HandleGetWithCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithCategory(itemID)
	if utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetWithGroup(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroup(itemID)
	if utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetWithGroupAndCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroupAndCategory(itemID)
	if utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundItem)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItems()
	if utils.HandleGetAllErrors(w, err, items, "items") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithCategories()
	if utils.HandleGetAllErrors(w, err, items, "items") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithGroups(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroups()
	if utils.HandleGetAllErrors(w, err, items, "items") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}

func (h *Handler) HandleGetAllWithGroupsAndCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroupsAndCategories()
	if utils.HandleGetAllErrors(w, err, items, "items") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, items)
}
