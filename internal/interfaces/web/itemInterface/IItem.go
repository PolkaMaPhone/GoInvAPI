package itemInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/itemDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
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
	r := chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGet)
	r.Get("/with_category", h.HandleGetWithCategory)
	r.Get("/with_group", h.HandleGetWithGroup)
	r.Get("/with_group_and_category", h.HandleGetWithGroupAndCategory)
	apiRouter.Mount(apiRouter.GetFullPath("/items/{item_id}"), r)

	r = chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/items", h.HandleGetAll)
	r.Get("/items_with_category", h.HandleGetAllWithCategories)
	r.Get("/items_with_group", h.HandleGetAllWithGroups)
	r.Get("/items_with_group_and_category", h.HandleGetAllWithGroupsAndCategories)
	apiRouter.Mount(apiRouter.GetFullPath("/"), r)
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
