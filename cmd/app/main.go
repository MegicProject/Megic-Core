package main

import (
	"Megic-core/config"
	"Megic-core/internal/db"
	"Megic-core/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.AppRoutes(e, db.GetDB(), cfg)

	port := ":8081"
	e.Logger.Fatal(e.Start(port))
}
