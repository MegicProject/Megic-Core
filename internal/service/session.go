package service

import (
	"math/rand"
	"time"
)

type SessionService interface {
	CreateSession() string
}

type sessionService struct{}

func NewSessionService() SessionService {
	return &sessionService{}
}

func (s *sessionService) CreateSession() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 50)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
