package handler

import (
	"Megic-core/config"
	"Megic-core/internal/model"
	"Megic-core/internal/service"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService    service.ChatService
	sessionService service.SessionService
	config         *config.Config
}

func NewChatHandler(chatService service.ChatService, sessionService service.SessionService, config *config.Config) *ChatHandler {
	return &ChatHandler{chatService: chatService, sessionService: sessionService, config: config}
}

func (h *ChatHandler) CreateSession(c echo.Context) error {
	var body struct {
		Message string `json:"message"`
	}

	if err := c.Bind(&body); err != nil {
		return ResponseBuilderInstance.InvalidRequest(c)
	}

	session := h.sessionService.CreateSession()

	chatPayload := model.ChatModel{
		Session: session,
		Message: body.Message,
		Role:    "user",
	}

	if err := h.chatService.CreateChat(chatPayload); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	return ResponseBuilderInstance.Success(c, session, "")
}

func (h *ChatHandler) GetChatsBySession(c echo.Context) error {
	session := c.Param("session")
	chats, err := h.chatService.GetChatsBySession(session)
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}
	if chats == nil {
		return ResponseBuilderInstance.Success(c, nil, "Error")
	}
	return ResponseBuilderInstance.Success(c, chats, "")
}

func (h *ChatHandler) CreateChat(c echo.Context) error {
	var chat model.ChatModel
	if err := c.Bind(&chat); err != nil {
		return ResponseBuilderInstance.InvalidRequest(c)
	}

	if err := h.chatService.CreateChat(chat); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	return ResponseBuilderInstance.Success(c, chat, "")
}
