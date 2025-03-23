package router

import (
	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
	"github.com/emersonvalentim/drobe-api/cmd/api/middlewares"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
	"github.com/emersonvalentim/drobe-api/services/auth"
	"github.com/emersonvalentim/drobe-api/services/inventory"
)

type Router struct {
	engine    *echo.Echo
	cfg       *config.Env
	inventory InventoryRouter
	auth      AuthRouter
}

func NewRouter(inventoryService *inventory.Service, authService *auth.Service, cfg *config.Env) *Router {
	e := echo.New()

	e.Use(middlewares.AuthMiddleware(authService))
	// Create API group for all routes
	api := e.Group("/api")

	return &Router{
		engine: e,
		cfg:    cfg,
		inventory: InventoryRouter{
			service: inventoryService,
			group:   api,
		},
		auth: AuthRouter{
			service: authService,
			group:   api,
		},
	}
}

func (r *Router) Register() {
	r.inventory.register(r.engine)
	r.auth.register(r.engine)
	r.listen(":" + r.cfg.Port)
}

func (r *Router) listen(port string) {
	r.engine.Logger.Fatal(r.engine.Start(port))
}

func getUserID(c echo.Context) uuid.UUID {
	userID := c.Get("userID")
	userIDUUID, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.UUID{}
	}
	return userIDUUID
}
