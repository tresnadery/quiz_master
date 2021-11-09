package dto

type RequestGetOrDeleteQuestion struct {
	Number string `json:"number" validate:"numeric"`
}
