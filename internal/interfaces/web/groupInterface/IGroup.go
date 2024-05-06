package groupInterface

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/domain/groupDomain"
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/validation"
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
	r := chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGet)
	apiRouter.Mount(apiRouter.GetFullPath("/groups/{group_id}"), r)

	r = chi.NewRouter()
	r.Use(validation.ValidateMethod(http.MethodGet))
	r.Get("/", h.HandleGetAll)
	apiRouter.Mount(apiRouter.GetFullPath("/groups"), r)
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	groupID, err := utils.GetIDFromRequest(w, r, idParameterName)
	if err != nil {
		return
	}

	foundGroup, err := h.service.GetGroupByID(groupID)
	if utils.HandleGetByIDErrors(w, err, foundGroup, groupID, "group") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, foundGroup)
}

func (h *Handler) HandleGetAll(w http.ResponseWriter, _ *http.Request) {
	groups, err := h.service.GetAllGroups()
	if utils.HandleGetAllErrors(w, err, groups, "groups") {
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, groups)
}
