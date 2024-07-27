package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/util"
	"time"

	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/google/uuid"
)

func UserFindAll() ([]entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	var users []entity.User
	db.Find(&users)
	return users, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserFindById(id string) (entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	var user entity.User
	db.Find(&user, "id = ?", id)
	if user.Id == uuid.Nil {
		return user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "user not found for id: " + id,
		}
	}
	return user, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserCreate(user *entity.User) (*entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	user.Id = uuid.New()
	passwordHash, err := util.HashPassword(user.Password)
	if err != nil {
		return user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: "password cannot be hashed",
			StackTrace:   err.Error(),
		}
	}
	user.Password = passwordHash
	user.CreatedAt = time.Now()
	err = db.Create(&user).Error
	if err != nil {
		return user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: "cannot create user",
			StackTrace:   err.Error(),
		}
	}
	return user, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserUpdate(id string, request *entity.User) (*entity.User, dto.ApplicationResponse) {
	user, _ := UserFindById(id)

	db := database.Database.Db
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.UpdatedAt = time.Now()
	err := db.Save(&user).Error
	if err != nil {
		return &user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: "cannot save user",
			StackTrace:   err.Error(),
		}
	}
	return &user, dto.ApplicationResponse{
		IsError: false,
	}
}

func userFindByUsername(username string) (*entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	var user entity.User
	db.Find(&user, "username = ?", username)
	if user.Id == uuid.Nil {
		return &user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "user not found for username: " + username,
		}
	}
	return &user, dto.ApplicationResponse{
		IsError: false,
	}
}
