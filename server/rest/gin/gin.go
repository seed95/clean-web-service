package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	"github.com/seed95/clean-web-service/config"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
	"github.com/seed95/clean-web-service/server/rest"
	"github.com/seed95/clean-web-service/service"
)

type (
	server struct {
		cfg         *config.Gin
		userService service.UserService
		logger      log.Logger
		translator  translate.Translator
	}

	Fields struct {
		Cfg         *config.Gin
		UserService service.UserService
		Logger      log.Logger
		Translator  translate.Translator
	}
)

var r *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()
}

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
	fmt.Printf("Gin http server started on %v:%v\n", aurora.Green(h.cfg.Host), aurora.Green(h.cfg.Port))
	return r.Run(fmt.Sprintf("%s:%d", h.cfg.Host, h.cfg.Port))
}

func (h *server) setRoutes() {
	r.POST("/user", h.createUser)
	r.GET("/user/:id", h.getUser)
	r.DELETE("/user/:id", h.deleteUser)
}
