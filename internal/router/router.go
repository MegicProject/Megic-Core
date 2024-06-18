package router

import (
	"database/sql"
	"warunggpt-core-service/internal/handler"
	"warunggpt-core-service/internal/repository"
	"warunggpt-core-service/internal/service"

	"github.com/labstack/echo/v4"
)

func AppRoutes(e *echo.Echo, db *sql.DB) {
	chatRepository := repository.NewChatRepository(db)
	chatService := service.NewChatService(chatRepository)
	chatHandler := handler.NewChatHandler(chatService)
	sessionService := service.NewSessionService()
	sessionHandler := handler.NewSessionHandler(sessionService)
	e.GET("/chats/:session", chatHandler.GetChatsBySession)
	e.POST("/chats", chatHandler.CreateChat)
	e.POST("/sessions", sessionHandler.CreateSession)
}
