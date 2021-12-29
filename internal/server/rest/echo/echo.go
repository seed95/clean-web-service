package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/internal/server/rest"
	"github.com/seed95/clean-web-service/internal/service"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
)

type (
	server struct {
		cfg         *config.Echo
		userService service.UserService
		logger      log.Logger
		translator  translate.Translator
	}

	Fields struct {
		Cfg         *config.Echo
		UserService service.UserService
		Logger      log.Logger
		Translator  translate.Translator
	}
)

var e = echo.New()

func New(fields *Fields) rest.Server {
	return &server{
		cfg:         fields.Cfg,
		userService: fields.UserService,
		logger:      fields.Logger,
		translator:  fields.Translator,
	}
}

func (h *server) Start() error {
	h.setRoutes()
	return e.Start(fmt.Sprintf(":%d", h.cfg.Port))
}

func (h *server) setRoutes() {
	e.POST("/user", h.createUser)
	e.GET("/user/:id", h.getUser)
	e.DELETE("/user/:id", h.deleteUser)
}
