package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/db"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
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
			r.Use(validationMiddleware.ValidateMethod(http.MethodGet, http.MethodPut, http.MethodDelete))
			r.Get("/", h.HandleGet)
			r.Get("/with_category", h.HandleGetWithCategory)
			r.Get("/with_group", h.HandleGetWithGroup)
			r.Get("/with_group_and_category", h.HandleGetWithGroupAndCategory)
			r.Delete("/", h.HandleDelete)
			r.Put("/", h.HandlePut)
		})

		r.With(validationMiddleware.ValidateMethod(http.MethodPost)).Post("/", h.HandlePost)

		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/", h.HandleGetAll)
		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/with_category", h.HandleGetAllWithCategories)
		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/with_group", h.HandleGetAllWithGroups)
		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/with_group_and_category", h.HandleGetAllWithGroupsAndCategories)
	})
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	logging.InfoLogger.Printf("Retrieved itemID: %v", itemID)

	foundItem, err := h.service.GetItemByID(itemID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundItem)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var item itemDomain.Item
	err := utils.DecodeFromRequest(w, r, &item)
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
	err = utils.RespondWithJSON(w, http.StatusCreated, createdItem)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
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
	err = utils.RespondWithJSON(w, http.StatusOK, response)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandlePut(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	existingItem, err := h.service.GetItemByID(itemID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, existingItem, itemID, "item")
		return
	}

	var newItem itemDomain.Item
	err = utils.DecodeFromRequest(w, r, &newItem)
	if err != nil {
		logging.ErrorLogger.Printf("Error decoding item from request: %v", err)
		return
	}
	logging.InfoLogger.Printf("New item: %v", newItem.ItemID)

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

	err = utils.RespondWithJSON(w, http.StatusOK, updatedItem)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetWithCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	foundItem, err := h.service.GetItemByIDWithCategory(itemID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundItem)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetWithGroup(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroup(itemID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundItem)
	if err != nil {
		return
	}
}

func (h *Handler) HandleGetWithGroupAndCategory(w http.ResponseWriter, r *http.Request) {
	itemID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	foundItem, err := h.service.GetItemByIDWithGroupAndCategory(itemID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundItem, itemID, "item")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundItem)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItems()
	if err != nil {
		utils.HandleGetAllErrors(w, err, items, "items")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, items)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAllWithCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithCategories()
	if err != nil {
		utils.HandleGetAllErrors(w, err, items, "items")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, items)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAllWithGroups(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroups()
	if err != nil {
		utils.HandleGetAllErrors(w, err, items, "items")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, items)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAllWithGroupsAndCategories(w http.ResponseWriter, _ *http.Request) {
	items, err := h.service.GetAllItemsWithGroupsAndCategories()
	if err != nil {
		utils.HandleGetAllErrors(w, err, items, "items")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, items)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}
