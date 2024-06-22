package handler

import (
	"Megic-core/config"
	"Megic-core/internal/model"
	"Megic-core/internal/service"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	chatService          service.ChatService
	config               *config.Config
	configurationService service.ConfigurationService
}

func NewChatHandler(chatService service.ChatService, config *config.Config, configurationService service.ConfigurationService) *ChatHandler {
	return &ChatHandler{chatService: chatService, config: config, configurationService: configurationService}
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

	instructionConfig, err := h.configurationService.GetByCode("AGENT_INSTRUCTION")
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	historyConfig, err := h.configurationService.GetByCode("AGENT_HISTORY")
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	var history []model.HistoryItem
	if err := json.Unmarshal([]byte(historyConfig.Value), &history); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	agentPayload := struct {
		Message     string              `json:"message"`
		Instruction string              `json:"instruction"`
		History     []model.HistoryItem `json:"history"`
	}{
		Message:     chat.Message,
		Instruction: instructionConfig.Value,
		History:     history,
	}

	marshalAgentPayload, err := json.Marshal(agentPayload)
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	agentApiUrl := h.config.AGENT_API_URL + "/chat"

	agentResponse, err := http.Post(agentApiUrl, "application/json", bytes.NewBuffer(marshalAgentPayload))
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}
	defer agentResponse.Body.Close()

	agentResponseBody, err := io.ReadAll(agentResponse.Body)
	if err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	type AgentResponseDto struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Message string `json:"message"`
		} `json:"data"`
	}

	var agentApiResponse AgentResponseDto
	if err := json.Unmarshal(agentResponseBody, &agentApiResponse); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	if agentApiResponse.Status != 200 {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	agentResponseChatData := model.ChatModel{
		Session: chat.Session,
		Message: agentApiResponse.Data.Message,
		Role:    "model",
	}

	if err := h.chatService.CreateChat(chat); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	if err := h.chatService.CreateChat(agentResponseChatData); err != nil {
		return ResponseBuilderInstance.InternalServiceError(c)
	}

	return ResponseBuilderInstance.Success(c, chat)
}
