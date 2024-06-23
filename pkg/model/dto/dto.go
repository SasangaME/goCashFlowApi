package dto

import "github.com/google/uuid"

type BuildDto struct {
	Version string `json:"version"`
	Env     string `json:"env"`
	Build   string `json:"buid"`
}

type RoleDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
