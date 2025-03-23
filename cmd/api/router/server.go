package router

import (
	"github.com/emersonvalentim/drobe-api/services/clothing"
	"github.com/labstack/echo/v4"
)

type Router struct {
	engine *echo.Echo

	cloth ClothRouter
}

func NewRouter(clothService *clothing.Service) *Router {
	e := echo.New()
	// Create API group for all routes
	api := e.Group("/api")

	return &Router{
		engine: e,
		cloth: ClothRouter{
			service: clothService,
			group:   api, // Pass the api group to ClothRouter
		},
	}
}

func (r *Router) Register() {
	r.cloth.register(r.engine)
	r.listen(":8080")
}

func (r *Router) listen(port string) {
	r.engine.Logger.Fatal(r.engine.Start(port))
}

type Reason struct {
	Reason string `json:"reason"`
}

func WithReason(reason string) Reason {
	return Reason{
		Reason: reason,
	}
}
