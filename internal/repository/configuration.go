package repository

import (
	"Megic-core/internal/model"
	"database/sql"
)

type ConfigurationRepository interface {
	Get() ([]model.ConfigurationModel, error)
	GetByCode(code string) (model.ConfigurationModel, error)
}

type configurationRepository struct {
	db *sql.DB
}

func NewConfigurationRepository(db *sql.DB) ConfigurationRepository {
	return &configurationRepository{db: db}
}

func (r *configurationRepository) Get() ([]model.ConfigurationModel, error) {
	query := "SELECT code, value FROM configurations"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configurations []model.ConfigurationModel
	for rows.Next() {
		var configuration model.ConfigurationModel
		if err := rows.Scan(&configuration.Code, &configuration.Value); err != nil {
			return nil, err
		}
		configurations = append(configurations, configuration)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return configurations, nil
}

func (r *configurationRepository) GetByCode(code string) (model.ConfigurationModel, error) {
	query := "SELECT code, value FROM configurations WHERE code = ?"
	row := r.db.QueryRow(query, code)

	var configuration model.ConfigurationModel
	if err := row.Scan(&configuration.Code, &configuration.Value); err != nil {
		return configuration, err
	}

	return configuration, nil
}
