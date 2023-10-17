package model

type ApiError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
