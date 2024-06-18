package handler

import (
	"warunggpt-core-service/internal/model"
	"warunggpt-core-service/internal/service"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (h *ChatHandler) GetChatsBySession(c echo.Context) error {
	session := c.Param("session")
	chats, err := h.chatService.GetChatsBySession(session)
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}
	return ResponseBuilderInstance.Success(c, chats)
}

func (h *ChatHandler) CreateChat(c echo.Context) error {
	var chat model.ChatModel
	if err := c.Bind(&chat); err != nil {
		return ResponseBuilderInstance.InvalidRequest(c)
	}

	if err := h.chatService.CreateChat(chat); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}
	return ResponseBuilderInstance.Success(c, chat)
}
