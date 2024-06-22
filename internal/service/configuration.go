package service

import (
	"Megic-core/internal/model"
	"Megic-core/internal/repository"
)

type ConfigurationService interface {
	Get() ([]model.ConfigurationModel, error)
	GetByCode(code string) (model.ConfigurationModel, error)
}

type configurationService struct {
	configurationRepository repository.ConfigurationRepository
}

func NewConfigurationService(configurationRepo repository.ConfigurationRepository) ConfigurationService {
	return &configurationService{configurationRepository: configurationRepo}
}

func (s *configurationService) Get() ([]model.ConfigurationModel, error) {
	return s.configurationRepository.Get()
}

func (s *configurationService) GetByCode(code string) (model.ConfigurationModel, error) {
	return s.configurationRepository.GetByCode(code)
}
