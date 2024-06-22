package router

import (
	"Megic-core/config"
	"Megic-core/internal/handler"
	"Megic-core/internal/repository"
	"Megic-core/internal/service"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func AppRoutes(e *echo.Echo, db *sql.DB, cfg *config.Config) {
	chatRepository := repository.NewChatRepository(db)
	configurationRepository := repository.NewConfigurationRepository(db)
	chatService := service.NewChatService(chatRepository)
	chatHandler := handler.NewChatHandler(chatService, cfg, configurationRepository)
	sessionService := service.NewSessionService()
	sessionHandler := handler.NewSessionHandler(sessionService)
	e.POST("/api/chats", chatHandler.CreateChat)
	e.GET("/api/chats/:session", chatHandler.GetChatsBySession)
	e.POST("/api/sessions", sessionHandler.CreateSession)
}
