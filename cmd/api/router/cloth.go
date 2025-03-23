package router

import (
	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
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

type ListClothResponse struct {
	Results []drobe.Cloth `json:"results"`
}

func (r *ClothRouter) ListClothes(c echo.Context) error {
	clothes, err := r.service.ListClothes(c.Request().Context())
	if err != nil {
		return c.JSON(500, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(200, ListClothResponse{
		Results: clothes,
	})
}

func (r *ClothRouter) GetCloth(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	cloth, err := r.service.GetCloth(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(404, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(200, cloth)
}

type CreateClothRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Color    string `json:"color"`
	Size     string `json:"size"`
	Brand    string `json:"brand"`
}

func (r *ClothRouter) CreateCloth(c echo.Context) error {
	input := CreateClothRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	cloth, err := r.service.CreateCloth(c.Request().Context(), clothing.CreateCloth{
		Name:     input.Name,
		Category: input.Category,
		Color:    input.Color,
		Size:     input.Size,
		Brand:    input.Brand,
	})
	if err != nil {
		return c.JSON(500, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(201, cloth)
}

func (r *ClothRouter) UpdateCloth(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (r *ClothRouter) DeleteCloth(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = r.service.DeleteCloth(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.NoContent(204)
}
