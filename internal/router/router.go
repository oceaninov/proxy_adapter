package router

import (
	"proxy-adapter/internal/handler"

	"github.com/labstack/echo/v4"
)

type Router struct {
	azureAd *azureAD
}

func NewRouter(server *echo.Echo, handlers handler.Handlers) (router *Router) {
	return &Router{
		azureAd: newAzureAD(server, handlers),
	}
}

func (r *Router) Initialize() {
	r.azureAd.initialize()
}
