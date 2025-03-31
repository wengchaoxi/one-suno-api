package repository

import (
	"github.com/wengchaoxi/one-suno-api/internal/model"
	"gorm.io/gorm"
)

type ApiKeyRepository struct {
	db *gorm.DB
}

func NewApiKeyRepository(db *gorm.DB) *ApiKeyRepository {
	return &ApiKeyRepository{db: db}
}

func (r *ApiKeyRepository) Create(apiKey *model.ApiKey) error {
	return r.db.Create(apiKey).Error
}

func (r *ApiKeyRepository) GetAll() ([]model.ApiKey, error) {
	var apiKeys []model.ApiKey
	err := r.db.Find(&apiKeys).Error
	if err != nil {
		return nil, err
	}
	return apiKeys, nil
}
