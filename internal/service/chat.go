package service

import (
	"Megic-core/internal/model"
	"Megic-core/internal/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatService interface {
	GetChatsBySession(session string) ([]model.ChatModel, error)
	CreateChat(chat model.ChatModel) error
}

type chatService struct {
	chatRepository       repository.ChatRepository
	configurationService ConfigurationService
	agentApiUrl          string
}

func NewChatService(chatRepo repository.ChatRepository, configurationService ConfigurationService, agentApiUrl string) ChatService {
	return &chatService{
		chatRepository:       chatRepo,
		configurationService: configurationService,
		agentApiUrl:          agentApiUrl,
	}
}

func (s *chatService) GetChatsBySession(session string) ([]model.ChatModel, error) {
	return s.chatRepository.Get(session)
}

func (s *chatService) CreateChat(chat model.ChatModel) error {
	instructionConfig, err := s.configurationService.GetByCode("AGENT_INSTRUCTION")

	if err != nil {
		return err
	}

	var messageHistory []model.HistoryItem

	history, err := s.chatRepository.Get(chat.Session)
	if err == nil {
		for _, h := range history {
			parts := []model.HistoryPart{
				{Text: h.Message},
			}
			messageHistory = append(messageHistory, model.HistoryItem{
				Parts: parts,
				Role:  h.Role,
			})
		}
	}

	agentPayload := struct {
		Message     string              `json:"message"`
		Instruction string              `json:"instruction"`
		History     []model.HistoryItem `json:"history"`
	}{
		Message:     chat.Message,
		Instruction: instructionConfig.Value,
		History:     messageHistory,
	}

	marshalAgentPayload, err := json.Marshal(agentPayload)
	if err != nil {
		return err
	}

	agentResponse, err := http.Post(s.agentApiUrl+"/chat", "application/json", bytes.NewBuffer(marshalAgentPayload))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer agentResponse.Body.Close()
	fmt.Println(agentResponse)

	agentResponseBody, err := io.ReadAll(agentResponse.Body)
	if err != nil {
		return err
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
		return err
	}

	if agentApiResponse.Status != 200 {
		return err
	}

	agentResponseChatData := model.ChatModel{
		Session: chat.Session,
		Message: agentApiResponse.Data.Message,
		Role:    "model",
	}

	if err := s.chatRepository.Create(chat); err != nil {
		return err
	}

	if err := s.chatRepository.Create(agentResponseChatData); err != nil {
		return err
	}

	return nil
}
