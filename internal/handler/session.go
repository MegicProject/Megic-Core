package handler

import (
	"fmt"
	"warunggpt-core-service/internal/service"

	"github.com/labstack/echo/v4"
)

type SessionHandler struct {
	service service.SessionService
}

func NewSessionHandler(sessionService service.SessionService) *SessionHandler {
	return &SessionHandler{service: sessionService}
}

type CreateSessionResponseDto struct {
	session string
}

func (h *SessionHandler) CreateSession(c echo.Context) error {
	session := h.service.CreateSession()
	fmt.Println(session)
	response := CreateSessionResponseDto{
		session: session,
	}
	return ResponseBuilderInstance.Success(c, response)
}
