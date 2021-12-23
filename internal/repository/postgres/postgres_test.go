package postgres

import (
	"github.com/seed95/clean-web-service/internal/config"
	"github.com/seed95/clean-web-service/internal/models"
	"github.com/seed95/clean-web-service/internal/models/types"
	"github.com/seed95/clean-web-service/internal/repository/postgres/schema"
	"github.com/seed95/clean-web-service/pkg/log/logrus"
	"github.com/seed95/clean-web-service/pkg/random"
	"github.com/seed95/clean-web-service/pkg/translate/i18n"
	"testing"
)

var (
	repositoryTest *repository
)

func setupTest(t *testing.T) {
	cfg := &config.Postgres{
		Username:  "postgres",
		Password:  "1233",
		DBName:    "clean-db-test",
		Host:      "localhost",
		Port:      "5432",
		SSLMode:   "disable",
		TimeZone:  "Asia/Tehran",
		Charset:   "utf8mb4",
		Migration: true,
	}

	opt := &logrus.Option{
		Path:         "../../../logs/test",
		Pattern:      "%Y-%m-%dT%H:%M",
		RotationSize: "20MB",
		RotationTime: "24h",
		MaxAge:       "720h",
	}
	logger, err := logrus.New(opt)
	if err != nil {
		t.Fatal(err)
	}

	translator, err := i18n.New("../../../build/i18n")
	if err != nil {
		t.Fatal(err)
	}

	repositoryTest = &repository{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}

	if err := repositoryTest.connect(); err != nil {
		t.Fatal(err)
	}

	if err := repositoryTest.db.Migrator().DropTable(&schema.User{}); err != nil {
		t.Fatal(err)
	}

	if err := repositoryTest.db.Migrator().CreateTable(&schema.User{}); err != nil {
		t.Fatal(err)
	}

}

func teardownTest() {
	repositoryTest = nil
}

func newUserTest() *models.User {
	return &models.User{
		Username:    random.String(9),
		Password:    random.String(20),
		Firstname:   random.String(9),
		Lastname:    random.String(9),
		Email:       random.String(7) + "@" + random.String(5) + "." + random.String(3),
		PhoneNumber: "0912" + random.StringWithCharset(7, "0123456789"),
		Role:        types.Basic,
		Gender:      types.Male,
	}
}
