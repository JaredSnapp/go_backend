package postgres

import (
	"fmt"
	"log"

	"github.com/JaredSnapp/go_backend/internal/config"
	"github.com/JaredSnapp/go_backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	conf *config.Config
	db   *gorm.DB
}

func NewService(conf *config.Config) *Service {
	return &Service{
		conf: conf,
	}
}

func (pg *Service) Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pg.conf.DBHost, pg.conf.DBPort, pg.conf.DBUser, pg.conf.DBPassword, pg.conf.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pg.db = db
	if err != nil {
		return err
	}

	return nil
}

func (pg *Service) Migrate() error {
	err := pg.db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = pg.db.AutoMigrate(&models.GoalMetaData{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
