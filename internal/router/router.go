package router

import (
	"database/sql"
	"warunggpt-core-service/config"
	"warunggpt-core-service/internal/handler"
	"warunggpt-core-service/internal/repository"
	"warunggpt-core-service/internal/service"

	"github.com/labstack/echo/v4"
)

func AppRoutes(e *echo.Echo, db *sql.DB, cfg *config.Config) {
	chatRepository := repository.NewChatRepository(db)
	configurationRepository := repository.NewConfigurationRepository(db)
	chatService := service.NewChatService(chatRepository)
	chatHandler := handler.NewChatHandler(chatService, cfg, configurationRepository)
	sessionService := service.NewSessionService()
	sessionHandler := handler.NewSessionHandler(sessionService)
	e.POST("/chats", chatHandler.CreateChat)
	e.GET("/chats/:session", chatHandler.GetChatsBySession)
	e.POST("/sessions", sessionHandler.CreateSession)
}
