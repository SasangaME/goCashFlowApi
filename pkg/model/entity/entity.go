package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Users       []User    `gorm:"foreignKey:RoleId"`
}

type User struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);null"`
	RoleId    uuid.UUID `gorm:"type:uuid;not null"`
}
