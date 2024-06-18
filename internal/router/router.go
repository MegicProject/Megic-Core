package router

import (
	"database/sql"
	"warunggpt-core-service/internal/handler"
	"warunggpt-core-service/internal/repository"
	"warunggpt-core-service/internal/service"

	"github.com/labstack/echo/v4"
)

func AppRoutes(e *echo.Echo, db *sql.DB) {
	// Chat routes
	chatRepository := repository.NewChatRepository(db)
	chatService := service.NewChatService(chatRepository)
	chatHandler := handler.NewChatHandler(chatService)
	e.GET("/chats/:session", chatHandler.GetChatsBySession)
	e.POST("/chats", chatHandler.CreateChat)
}
