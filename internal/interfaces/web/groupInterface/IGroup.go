package groupInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validationMiddleware"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

const idParameterName = "group_id"

type Handler struct {
	service *groupDomain.Service
}

func NewGroupHandler(s *groupDomain.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) HandleRoutes(apiRouter *customRouter.CustomRouter) {
	apiRouter.Route("/api/groups", func(r chi.Router) {
		r.Route("/{group_id}", func(r chi.Router) {
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
	groupID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	logging.InfoLogger.Printf("Retrieved groupID: %v", groupID)

	foundGroup, err := h.service.GetGroupByID(groupID)
	if err != nil {
		utils.HandleGetByIDErrors(w, err, foundGroup, groupID, "group")
		return
	}

	err = utils.RespondWithJSON(w, http.StatusOK, foundGroup)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	groups, err := h.service.GetAllGroups()
	if err != nil {
		utils.HandleGetAllErrors(w, err, groups, "groups")
		return
	}
	logging.InfoLogger.Printf("Retrieved groups: %v", groups)

	err = utils.RespondWithJSON(w, http.StatusOK, groups)
	if err != nil {
		utils.HandleRespondWithJSONErrors(w, err)
		return
	}
}

func (h *Handler) HandleDelete(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Group Delete not implemented yet")
	http.Error(writer, "Not implemented", http.StatusNotImplemented)
	//TODO - Implement Group Delete
}

func (h *Handler) HandlePut(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Group update not implemented yet")
	http.Error(writer, "Not implemented", http.StatusNotImplemented)
	//TODO - Implement Group update
}

func (h *Handler) HandlePost(writer http.ResponseWriter, request *http.Request) {
	logging.ErrorLogger.Printf("Group create not implemented yet")
	http.Error(writer, "Not implemented", http.StatusNotImplemented)
	//TODO - Implement Group create
}
