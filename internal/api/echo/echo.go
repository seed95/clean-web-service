package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/seed95/clean-web-service/internal/api"
	"github.com/seed95/clean-web-service/internal/service"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
)

type (
	httpServer struct {
		userService service.UserService
		logger      log.Logger
		translator  translate.Translator
	}

	Fields struct {
		UserService service.UserService
		Logger      log.Logger
		Translator  translate.Translator
	}
)

var e = echo.New()

func New(fields *Fields) api.HttpServer {
	return &httpServer{
		userService: fields.UserService,
		logger:      fields.Logger,
		translator:  fields.Translator,
	}
}

func (h *httpServer) Start(port int) error {
	h.setRoutes()
	return e.Start(fmt.Sprintf(":%d", port))
}

func (h *httpServer) setRoutes() {
	e.POST("/user", h.createUser)
	e.GET("/user/:id", h.getUser)
	e.DELETE("/user/:id", h.deleteUser)
}
