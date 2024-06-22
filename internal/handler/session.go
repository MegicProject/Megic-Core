package handler

import (
	"Megic-core/internal/service"

	"github.com/labstack/echo/v4"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(sessionService service.SessionService) *SessionHandler {
	return &SessionHandler{service: sessionService}
}

type CreateSessionResponseDto struct {
	Session string `json:"session"`
}

func (h *SessionHandler) CreateSession(c echo.Context) error {
	session := h.service.CreateSession()
	response := CreateSessionResponseDto{
		Session: session,
	}
	return ResponseBuilderInstance.Success(c, response)
}
