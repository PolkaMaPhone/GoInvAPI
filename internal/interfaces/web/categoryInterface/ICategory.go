package categoryInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/categoryDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
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
	apiRouter.Route("/api/categories", func(r chi.Router) {
		r.Route("/{category_id}", func(r chi.Router) {
			r.Use(validationMiddleware.ValidateMethod(http.MethodGet, http.MethodPut, http.MethodDelete))
			r.Get("/", h.HandleGet)
			r.Delete("/", h.HandleDelete)
			r.Put("/", h.HandlePut)
		})

		r.With(validationMiddleware.ValidateMethod(http.MethodPost)).Post("/", h.HandlePost)

		r.With(validationMiddleware.ValidateMethod(http.MethodGet)).Get("/", h.HandleGetAll)
	})
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	logging.InfoLogger.Printf("Retrieved categoryID: %v", categoryID)

	foundCategory, err := h.service.GetCategoryByID(categoryID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundCategory, categoryID, "category")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundCategory)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		utils.HandleGetAllErrors(w, err, categories, "categories")
		return
	}
	logging.InfoLogger.Printf("Retrieved %v categories.", len(categories))
	err = utils.RespondWithJSON(w, http.StatusOK, categories)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Delete method not implemented")
	http.Error(writer, "Delete method not implemented", http.StatusNotImplemented)
	//TODO - Implement delete category method
}

func (h *Handler) HandlePut(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Put method not implemented")
	http.Error(writer, "Put method not implemented", http.StatusNotImplemented)
	//TODO - Implement put category method
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	logging.ErrorLogger.Printf("Post method not implemented")
	http.Error(w, "Post method not implemented", http.StatusNotImplemented)
	//TODO - Implement post category method
}
