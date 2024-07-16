package database

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"gorm.io/gorm"
	"log"
)

func addMigrations(db *gorm.DB) {
	log.Println("running migrations")
	err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
	)
	if err != nil {
		return
	}

}
