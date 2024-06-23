package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
}

func (role *Role) BeforeCrateRole(tx *gorm.DB) (err error) {
	//UUID version 4
	role.Id = uuid.New()
	role.CreatedAt = time.Now()
	return
}
