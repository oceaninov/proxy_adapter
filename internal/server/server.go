package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	commons "proxy-adapter/internal/common"
	"proxy-adapter/internal/handler"
	"proxy-adapter/internal/router"
	"proxy-adapter/internal/service"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type IServer interface {
	StartApp()
}

type server struct {
	opt      commons.Options
	services *service.Services
}

// NewServer create object server
func NewServer(opt commons.Options, services *service.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func initHandler(opt commons.Options, services *service.Services) (handlers handler.Handlers) {
	hOpt := handler.HandlerOption{
		Options:  opt,
		Services: services,
	}
	return handler.Handlers{
		AzureAD: handler.AzureADHandler{
			HandlerOption: hOpt,
		},
	}
}

func (s *server) StartApp() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())

	handlers := initHandler(s.opt, s.services)
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		s.opt.Logger.Info("[API] Server is shutting down")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// We received an interrupt signal, shut down.
		if err := e.Shutdown(ctx); err != nil {
			s.opt.Logger.Error("[API] Fail to shutting down",
				zap.Error(err),
			)
		}
		close(idleConnectionClosed)
	}()

	routers := router.NewRouter(e, handlers)
	routers.Initialize()

	srvAddr := fmt.Sprintf("%s:%d", s.opt.Config.AppHost, cast.ToInt(s.opt.Config.AppPort))
	s.opt.Logger.Info(fmt.Sprintf("[API] HTTP serve at %s", srvAddr))
	if err := e.Start(srvAddr); err != nil {
		s.opt.Logger.Error("[API] Fail to start listen and server",
			zap.Error(err),
		)
	}

	<-idleConnectionClosed
	s.opt.Logger.Info("[API] Bye")
}
