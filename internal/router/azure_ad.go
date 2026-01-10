package router

import (
	"proxy-adapter/internal/handler"
	"proxy-adapter/internal/middleware"

	"github.com/labstack/echo/v4"
)

type azureAD struct {
	server   *echo.Echo
	handlers handler.Handlers
}

func newAzureAD(server *echo.Echo, handlers handler.Handlers) *azureAD {
	return &azureAD{
		server:   server,
		handlers: handlers,
	}
}

func (h *azureAD) initialize() {
	p := h.server.Group("/azure-ad")
	p.POST("/auth", middleware.HandlerWrapperJson(h.handlers.AzureAD.GetAccessToken))
}
