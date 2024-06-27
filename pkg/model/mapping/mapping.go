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
