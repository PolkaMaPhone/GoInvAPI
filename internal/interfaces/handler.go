package interfaces

import (
	"github.com/PolkaMaPhone/GoInvAPI/internal/infrastructure/customRouter"
)

type Handler interface {
	HandleRoutes(router *customRouter.CustomRouter)
}
