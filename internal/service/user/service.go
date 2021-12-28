package user

import (
	repo "github.com/seed95/clean-web-service/internal/repository"
	service "github.com/seed95/clean-web-service/internal/service"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
)

type (
	handler struct {
		userRepo   repo.UserRepository
		validation service.ValidationService
		logger     log.Logger
		translator translate.Translator
	}

	Option struct {
		UserRepo   repo.UserRepository
		Validation service.ValidationService
		Logger     log.Logger
		Translator translate.Translator
	}
)

func New(opt *Option) service.UserService {
	return &handler{
		userRepo:   opt.UserRepo,
		validation: opt.Validation,
		logger:     opt.Logger,
		translator: opt.Translator,
	}
}
