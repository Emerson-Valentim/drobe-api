package router

import (
	"net/http"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
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
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
	}

	user, err := r.service.SignUp(c.Request().Context(), auth.SignUp{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}

	return config.NewApiResponse(http.StatusCreated).WithData(user).Send(c)
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthRouter) SignIn(c echo.Context) error {
	request := SignInRequest{}
	if err := c.Bind(&request); err != nil {
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
	}

	token, err := r.service.SignIn(c.Request().Context(), auth.SignIn{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}

	return config.NewApiResponse(http.StatusOK).WithData(token).Send(c)
}
