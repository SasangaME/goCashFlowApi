package dto

type BuildDto struct {
	Version string `json:"version"`
	Env     string `json:"env"`
	Build   string `json:"build"`
}
