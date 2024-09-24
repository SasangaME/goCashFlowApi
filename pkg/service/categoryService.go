package service

import (
	"strconv"
	"time"

	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/SasangaME/goCashFlowApi/pkg/repository"
	"github.com/google/uuid"
)

func CategoryFindAll() ([]entity.Category, dto.ApplicationResponse) {
	categories, err := repository.CategoryFindAll()
	if err != nil {
		return nil, dto.ApplicationResponse{
			IsError:      true,
			ErrorMessage: "cannot find categories",
			StatusCode:   constants.InternalServerError,
			StackTrace:   err.Error(),
		}
	}
	return categories, dto.ApplicationResponse{
		IsError: false,
	}
}

func CategoryFindById(id int64) (entity.Category, dto.ApplicationResponse) {
	category, err := repository.CategoryFindById(id)
	if err != nil {
		return category, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: "unexpected error",
			StackTrace:   err.Error(),
		}
	}

	if category.Id == 0 {
		return category, dto.ApplicationResponse{
			IsError:      true,
			ErrorMessage: "category not found for id: " + strconv.FormatInt(id, 10),
			StatusCode:   constants.NotFound,
		}
	}

	return category, dto.ApplicationResponse{
		IsError: false,
	}
}

func CategoryCreate(category *entity.Category, userId uuid.UUID) {
	category.CreatedAt = time.Now()
	repository.CategoryCreate(*category)
}

func CategoryUpdate(category *entity.Category, userId uuid.UUID) {
	existingCategory, _ := CategoryFindById(category.Id)
	existingCategory.UpdatedAt = time.Now()
	repository.CategoryUpdate(existingCategory)
}
