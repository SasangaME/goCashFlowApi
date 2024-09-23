package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/util"
)

func Login(username, password string) (string, dto.ApplicationResponse) {
	user, err := UserFindByUsername(username)
	if err.IsError {
		return "", err
	}

	isAuthenticated := util.VerifyPassword(password, user.Password)
	if !isAuthenticated {
		return "", dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.NotAuthorized,
			ErrorMessage: "invalid password",
		}
	}

	jwt, jwtErr := util.CreateJWT(username)
	if jwtErr != nil {
		return "", dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: "jwt creation error",
			StackTrace:   jwtErr.Error(),
		}
	}
	return jwt, err
}

func Authorize(jwt string, roles []string) (int64, dto.ApplicationResponse) {
	username, err := util.ValidateJWT(jwt)
	if err != nil {
		return 0, dto.ApplicationResponse{
			IsError:    true,
			StatusCode: constants.NotAuthorized,
		}
	}

	user, userResponse := UserFindByUsername(username)
	if userResponse.IsError {
		return 0, userResponse
	}

	isValid := false
	for _, role := range roles {
		if user.Role.Name == role {
			isValid = true
			break
		}
	}

	if !isValid {
		return 0, dto.ApplicationResponse{
			IsError:    true,
			StatusCode: constants.NotAuthorized,
		}
	}

	return user.Id, dto.ApplicationResponse{
		IsError: false,
	}
}
