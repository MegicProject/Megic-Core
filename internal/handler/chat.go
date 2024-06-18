package handler

import (
	"net/http"
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
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, chats)
}

func (h *ChatHandler) CreateChat(c echo.Context) error {
	var chat model.ChatModel
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.chatService.CreateChat(chat); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, chat)
}
