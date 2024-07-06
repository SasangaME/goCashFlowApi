package mapping

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
)

func RoleToDto(role entity.Role) *dto.RoleDto {
	return &dto.RoleDto{
		Id:          role.Id,
		Name:        role.Name,
		Description: role.Description,
	}
}

func DtoToRole(dto *dto.RoleDto) entity.Role {
	return entity.Role{
		Name:        dto.Name,
		Description: dto.Description,
	}
}

func RoleListToDto(roles []entity.Role) []*dto.RoleDto {
	roleList := make([]*dto.RoleDto, len(roles))
	for i, role := range roles {
		roleList[i] = RoleToDto(role)
	}
	return roleList
}

func UserRequestToUser(request *dto.UserRequest) entity.User {
	return entity.User{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		RoleId:    request.RoleId,
		Email:     request.Email,
	}
}

func UserToUserResponse(user entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleId:    user.RoleId,
		Email:     user.Email,
	}
}

func UserListToUserResponse(users []entity.User) []*dto.UserResponse {
	userList := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		userList[i] = UserToUserResponse(user)
	}
	return userList
}
