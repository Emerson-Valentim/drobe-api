package router

import (
	"net/http"

	"github.com/emersonvalentim/drobe-api/services/auth"
	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	service *auth.Service
	group   *echo.Group
}

func (r *AuthRouter) register(e *echo.Echo) {
	auth := r.group.Group("/auth")
	auth.POST("/signup", r.SignUp)
	auth.POST("/signin", r.SignIn)
}

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r *AuthRouter) SignUp(c echo.Context) error {
	request := SignUpRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	user, err := r.service.SignUp(c.Request().Context(), auth.SignUp{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRouter) SignIn(c echo.Context) error {
	request := SignInRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	token, err := r.service.SignIn(c.Request().Context(), auth.SignIn{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, token)
}
