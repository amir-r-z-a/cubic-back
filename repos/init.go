package repos

import (
	"github.com/amir-r-z-a/cubic-back/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppRepo struct {
	DB *gorm.DB
}

func InitRepo(appcfg *config.AppConfig) *AppRepo {

	db := newPostgresDB(appcfg.PostgresDSN)

	return &AppRepo{
		DB: db,
	}

}

func newPostgresDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
