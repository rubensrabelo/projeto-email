package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=golang_db port=5555 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db, nil
}
