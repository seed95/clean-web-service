package postgres

import (
	"fmt"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/config"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/pkg/translate"
	repo "github.com/seed95/clean-web-service/repository"
	"github.com/seed95/clean-web-service/repository/postgres/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	repository struct {
		db         *gorm.DB
		cfg        *config.Postgres
		logger     log.Logger
		translator translate.Translator
	}
)

func New(cfg *config.Postgres, logger log.Logger, translator translate.Translator) (repo.Repository, error) {
	postgresRepository := &repository{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}

	if err := postgresRepository.connect(); err != nil {
		return nil, err
	}

	if cfg.Migration {
		if err := postgresRepository.migration(); err != nil {
			return nil, err
		}
	}

	return postgresRepository, nil
}

func (r *repository) connect() error {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		r.cfg.Host,
		r.cfg.Username,
		r.cfg.Password,
		r.cfg.DBName,
		r.cfg.Port,
		r.cfg.SSLMode,
		r.cfg.TimeZone)
	postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disable default gorm log
	})
	if err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres",
			Function: "connect",
			Message:  err.Error(),
		})

		return derrors.New(derrors.Unexpected, messages.DBError)
	}

	r.db = postgresDB

	return nil
}

func (r *repository) migration() error {

	if err := r.db.AutoMigrate(&schema.User{}); err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres",
			Function: "migration",
			Message:  err.Error(),
		})
		return derrors.New(derrors.Unexpected, messages.DBError)
	}

	return nil

}
