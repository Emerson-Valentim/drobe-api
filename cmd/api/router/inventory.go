package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/emersonvalentim/drobe-api/cmd/api/config"
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
	items.POST("/:id/image", r.UploadImage)
}

func (r *InventoryRouter) ListItems(c echo.Context) error {
	items, err := r.service.ListItems(c.Request().Context(), getUserID(c))
	if err != nil {
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}
	return config.NewApiResponse(http.StatusOK).WithData(items).Send(c)
}

func (r *InventoryRouter) GetItem(c echo.Context) error {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
	}

	item, err := r.service.GetItem(c.Request().Context(), itemID, getUserID(c))
	if err != nil {
		return config.NewApiResponse(http.StatusNotFound).WithMessage(err.Error()).Send(c)
	}
	return config.NewApiResponse(http.StatusOK).WithData(item).Send(c)
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
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
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
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}
	return config.NewApiResponse(http.StatusCreated).WithData(item).Send(c)
}

func (r *InventoryRouter) UpdateItem(c echo.Context) error {
	return nil
}

func (r *InventoryRouter) DeleteItem(c echo.Context) error {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
	}

	err = r.service.DeleteItem(c.Request().Context(), itemID, getUserID(c))
	if err != nil {
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}

	return c.NoContent(http.StatusNoContent)
}

func (r *InventoryRouter) UploadImage(c echo.Context) error {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return config.NewApiResponse(http.StatusBadRequest).WithMessage(err.Error()).Send(c)
	}

	uploadURL, err := r.service.UploadImage(c.Request().Context(), itemID, getUserID(c))
	if err != nil {
		return config.NewApiResponse(http.StatusInternalServerError).WithMessage(err.Error()).Send(c)
	}

	return config.NewApiResponse(http.StatusOK).WithData(map[string]string{
		"uploadURL": uploadURL,
	}).Send(c)
}
