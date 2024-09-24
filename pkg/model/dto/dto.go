package dto

import "github.com/google/uuid"

type BuildDto struct {
	Version string `json:"version"`
	Env     string `json:"env"`
	Build   string `json:"build"`
}

type RoleDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserRequest struct {
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    uuid.UUID `json:"roleId"`
}

type UserResponse struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	RoleId    uuid.UUID `json:"role_id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CategoryDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
