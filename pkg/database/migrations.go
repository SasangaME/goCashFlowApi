package database

import (
	"log"

	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"gorm.io/gorm"
)

func addMigrations(db *gorm.DB) {
	log.Println("running migrations")
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Category{},
	)
	if err != nil {
		return
	}

}
