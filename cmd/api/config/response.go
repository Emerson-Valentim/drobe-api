package config

import (
	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Status  int         `json:"-"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewApiResponse(status int) *ApiResponse {
	return &ApiResponse{
		Status: status,
	}
}

func (r *ApiResponse) WithMessage(message string) *ApiResponse {
	r.Message = message
	return r
}

func (r *ApiResponse) WithStatus(status int) *ApiResponse {
	r.Status = status
	return r
}

func (r *ApiResponse) WithData(data interface{}) *ApiResponse {
	r.Data = data
	return r
}

func (r *ApiResponse) Send(c echo.Context) error {
	return c.JSON(r.Status, r)
}
