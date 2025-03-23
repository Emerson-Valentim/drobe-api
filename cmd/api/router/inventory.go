package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
	"github.com/emersonvalentim/drobe-api/services/inventory"
)

type InventoryRouter struct {
	service *inventory.Service
	group   *echo.Group
}

func (r *InventoryRouter) register(e *echo.Echo) {
	items := r.group.Group("/inventory")
	items.GET("", r.ListItems)
	items.GET("/:id", r.GetItem)
	items.POST("", r.CreateItem)
	items.PUT("/:id", r.UpdateItem)
	items.DELETE("/:id", r.DeleteItem)
}

type ListInventoryResponse struct {
	Results []drobe.Item `json:"results"`
}

func (r *InventoryRouter) ListItems(c echo.Context) error {
	items, err := r.service.ListItems(c.Request().Context(), getUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ListInventoryResponse{
		Results: items,
	})
}

func (r *InventoryRouter) GetItem(c echo.Context) error {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	item, err := r.service.GetItem(c.Request().Context(), itemID, getUserID(c))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, item)
}

type CreateInventoryRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Color    string `json:"color"`
	Size     string `json:"size"`
	Brand    string `json:"brand"`
}

func (r *InventoryRouter) CreateItem(c echo.Context) error {
	input := CreateInventoryRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	item, err := r.service.CreateItem(c.Request().Context(), inventory.CreateInventoryItem{
		Name:     input.Name,
		Category: input.Category,
		Color:    input.Color,
		Size:     input.Size,
		Brand:    input.Brand,
		OwnerID:  getUserID(c),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, item)
}

func (r *InventoryRouter) UpdateItem(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (r *InventoryRouter) DeleteItem(c echo.Context) error {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = r.service.DeleteItem(c.Request().Context(), itemID, getUserID(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
