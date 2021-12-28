package application

import (
	"github.com/seed95/clean-web-service/internal/api/echo"
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/internal/repository/postgres"
	"github.com/seed95/clean-web-service/internal/service/user"
	"github.com/seed95/clean-web-service/internal/service/validation"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/log/logrus"
	"github.com/seed95/clean-web-service/pkg/translate/i18n"
)

func Run(cfg *config.Config) error {
	logger, err := logrus.New(&logrus.Option{
		Path:         cfg.Logger.Logrus.Path,
		Pattern:      cfg.Logger.Logrus.Pattern,
		RotationSize: cfg.Logger.Logrus.RotationSize,
		RotationTime: cfg.Logger.Logrus.RotationTime,
		MaxAge:       cfg.Logger.Logrus.MaxAge,
	})
	if err != nil {
		return err
	}

	translator, err := i18n.New(cfg.Translator.I18n.MessagePath)
	if err != nil {
		logger.Error(&log.Field{
			Section:  "application",
			Function: "Run",
			Message:  err.Error(),
		})
	}

	repo, err := postgres.New(&cfg.Database.Postgres, logger, translator)
	if err != nil {
		return err
	}

	_ = repo

	validationService := validation.New(&cfg.Validation, logger, translator)
	userService := user.New(&user.Option{
		UserRepo:   repo,
		Validation: validationService,
		Logger:     logger,
		Translator: translator,
	})

	server := echo.New(&echo.Fields{
		UserService: userService,
		Logger:      logger,
		Translator:  translator,
	})

	return server.Start(8085)
}
