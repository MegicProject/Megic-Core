package service

import (
	"Megic-core/internal/model"
	"Megic-core/internal/repository"
)

type ChatService interface {
	GetChatsBySession(session string) ([]model.ChatModel, error)
	CreateChat(chat model.ChatModel) error
}

type chatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepo repository.ChatRepository) ChatService {
	return &chatService{chatRepository: chatRepo}
}

func (s *chatService) GetChatsBySession(session string) ([]model.ChatModel, error) {
	return s.chatRepository.Get(session)
}

func (s *chatService) CreateChat(chat model.ChatModel) error {
	return s.chatRepository.Create(chat)
}
