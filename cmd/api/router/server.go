package router

import (
	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/services/clothing"
)

type Router struct {
	engine *echo.Echo
	cfg    *config.Env
	cloth  ClothRouter
}

func NewRouter(clothService *clothing.Service, cfg *config.Env) *Router {
	e := echo.New()
	// Create API group for all routes
	api := e.Group("/api")

	return &Router{
		engine: e,
		cfg:    cfg,
		cloth: ClothRouter{
			service: clothService,
			group:   api, // Pass the api group to ClothRouter
		},
	}
}

func (r *Router) Register() {
	r.cloth.register(r.engine)
	r.listen(":" + r.cfg.Port)
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
