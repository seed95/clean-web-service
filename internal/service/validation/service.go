package validation

import (
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/internal/service"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
)

type (
	handler struct {
		cfg        *config.Validation
		logger     log.Logger
		translator translate.Translator
	}
)

func New(cfg *config.Validation, logger log.Logger, translator translate.Translator) service.ValidationService {
	return &handler{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}
}
