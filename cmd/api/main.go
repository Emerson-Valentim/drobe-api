package main

import (
	"github.com/emersonvalentim/drobe-api/cmd/api/router"
	"github.com/emersonvalentim/drobe-api/services/clothing"
)

func main() {
	clothService := clothing.NewService()

	router := router.NewRouter(clothService)
	router.Register()
}
