package repository

import (
	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
)

func CategoryFindAll() ([]entity.Category, error) {
	db := database.Database.Db
	var categories []entity.Category
	db.Find(&categories)
	return categories, nil
}

func CategoryFindById(id int64) (entity.Category, error) {
	db := database.Database.Db
	var category entity.Category
	db.Find(&category, "id=?", id)
	return category, nil
}

func CategoryCreate(category entity.Category) error {
	db := database.Database.Db
	err := db.Create(category).Error
	return err
}

func CategoryUpdate(category entity.Category) error {
	db := database.Database.Db
	err := db.Save(category).Error
	return err
}
