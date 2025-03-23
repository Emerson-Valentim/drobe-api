package router

import (
	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api/services/clothing"
)

type ClothRouter struct {
	service *clothing.Service
	group   *echo.Group
}

func (r *ClothRouter) register(e *echo.Echo) {
	clothes := r.group.Group("/clothes")
	clothes.GET("", r.ListClothes)
	clothes.GET("/:id", r.GetCloth)
	clothes.POST("", r.CreateCloth)
	clothes.PUT("/:id", r.UpdateCloth)
	clothes.DELETE("/:id", r.DeleteCloth)
}

func (r *ClothRouter) ListClothes(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (r *ClothRouter) GetCloth(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (r *ClothRouter) CreateCloth(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (r *ClothRouter) UpdateCloth(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (r *ClothRouter) DeleteCloth(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}
