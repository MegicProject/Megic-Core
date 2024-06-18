package main

import (
	"warunggpt-core-service/config"
	"warunggpt-core-service/internal/db"
	"warunggpt-core-service/internal/router"

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

	router.AppRoutes(e, db.GetDB().DB)

	e.Logger.Fatal(e.Start(":8080"))
}
