package entity

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedBy uint
	CreatedAt time.Time
	UpdatedBY uint
	UpdatedAt time.Time
}

type Role struct {
	BaseEntity
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Users       []User `gorm:"foreignKey:RoleId"`
}

type User struct {
	BaseEntity
	Username  string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);null"`
	RoleId    uuid.UUID `gorm:"type:uuid;not null"`
	Role      Role      `gorm:"foreignKey:RoleId"`
}

type Category struct {
	BaseEntity
	Code        string `gorm:"type:varchar(255);not null;uniqueIndex:idx_category_code"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	// CreatedBy   uuid.UUID `gorm:"type:uuid; null"`
	// UpdatedBy   uuid.UUID `gorm:"type:uuid; null"`
}
