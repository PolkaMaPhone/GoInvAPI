package categoryInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validation"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const idParameterName = "category_id"

type Handler struct {
	service *categoryDomain.Service
}

func NewCategoryHandler(s *categoryDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(apiRouter *customRouter.CustomRouter) {
	r := chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGet)
	apiRouter.Mount(apiRouter.GetFullPath("/categories/{category_id}"), r)

	r = chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGetAll)
	apiRouter.Mount(apiRouter.GetFullPath("/categories"), r)
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundCategory, err := h.service.GetCategoryByID(categoryID)
	if utils.HandleGetByIDErrors(w, err, foundCategory, categoryID, "category") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundCategory)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	categories, err := h.service.GetAllCategories()
	if utils.HandleGetAllErrors(w, err, categories, "categories") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, categories)
}
