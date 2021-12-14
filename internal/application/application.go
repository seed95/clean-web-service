package application

import (
	"fmt"
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/log/logrus"
	"github.com/seed95/clean-web-service/pkg/translate"
	"github.com/seed95/clean-web-service/pkg/translate/i18n"
	"github.com/seed95/clean-web-service/pkg/translate/messages"
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

	for i := 0; i < 1; i++ {
		logger.Error(&log.Field{
			Section:  "application",
			Function: "Run",
			Params:   map[string]interface{}{"config": cfg},
			Message:  "Test logger",
		})
	}

	translator, err := i18n.New(cfg.Translator.I18N.MessagePath)
	if err != nil {
		logger.Error(&log.Field{
			Section:  "application",
			Function: "Run",
			Message:  err.Error(),
		})
	}
	lang := translate.EN
	fmt.Println(translator.Translate(messages.UserNotFound, lang))
	return nil
}
