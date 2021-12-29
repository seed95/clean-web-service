package application

import (
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/internal/repository/postgres"
	"github.com/seed95/clean-web-service/internal/server/rest"
	"github.com/seed95/clean-web-service/internal/server/rest/echo"
	"github.com/seed95/clean-web-service/internal/server/rest/gin"
	"github.com/seed95/clean-web-service/internal/service/user"
	"github.com/seed95/clean-web-service/internal/service/validation"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/log/logrus"
	"github.com/seed95/clean-web-service/pkg/translate"
	"github.com/seed95/clean-web-service/pkg/translate/i18n"
)

var cfg = &config.Config{}

type (
	Options struct {
		ConfigFile string
		RestType   string
	}
)

func Run(opt *Options) error {

	if err := initConfig(opt.ConfigFile); err != nil {
		return err
	}

	logger, err := initLog()
	if err != nil {
		return err
	}

	translator, err := initTranslator()
	if err != nil {
		logger.Error(&log.Field{
			Section:  "application",
			Function: "Run",
			Message:  err.Error(),
		})

		return err
	}

	repo, err := postgres.New(&cfg.Database.Postgres, logger, translator)
	if err != nil {
		return err
	}

	validationService := validation.New(&cfg.Validation, logger, translator)
	userService := user.New(&user.Option{
		UserRepo:   repo,
		Validation: validationService,
		Logger:     logger,
		Translator: translator,
	})

	var server rest.Server
	if opt.RestType == "echo" {
		server = echo.New(&echo.Fields{
			Cfg:         &cfg.Server.Rest.Echo,
			UserService: userService,
			Logger:      logger,
			Translator:  translator,
		})
	} else if opt.RestType == "gin" {
		server = gin.New(&gin.Fields{
			Cfg:         &cfg.Server.Rest.Gin,
			UserService: userService,
			Logger:      logger,
			Translator:  translator,
		})
	}

	return server.Start()
}

func initConfig(configFile string) error {
	return config.Parse(configFile, cfg)
}

func initLog() (log.Logger, error) {
	return logrus.New(&logrus.Option{
		Path:         cfg.Logger.Logrus.Path,
		Pattern:      cfg.Logger.Logrus.Pattern,
		RotationSize: cfg.Logger.Logrus.RotationSize,
		RotationTime: cfg.Logger.Logrus.RotationTime,
		MaxAge:       cfg.Logger.Logrus.MaxAge,
	})
}

func initTranslator() (translate.Translator, error) {
	return i18n.New(cfg.Translator.I18n.MessagePath)
}
