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
	sessionService := service.NewSessionService()
	chatService := service.NewChatService(chatRepository, configurationRepository, config.LoadConfig().AGENT_API_URL)
	chatHandler := handler.NewChatHandler(chatService, sessionService, cfg)
	e.POST("/api/chats", chatHandler.CreateChat)
	e.POST("/api/chats/session", chatHandler.CreateSession)
	e.GET("/api/chats/:session", chatHandler.GetChatsBySession)
}
