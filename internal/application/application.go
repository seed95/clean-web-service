package application

import (
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/log/logrus"
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
	return nil
}
