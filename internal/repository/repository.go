package repository

import "github.com/wengchaoxi/one-suno-api/internal/model"

type UserRepositoryInterface interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	// Update(id string, updates map[string]any) error
	// Delete(id string) error
}

type ApiKeyRepositoryInterface interface {
	Create(apiKey *model.ApiKey) error
	GetAll() ([]model.ApiKey, error)
	// Update(id string, updates map[string]any) error
	// Delete(id string) error
}

type ProviderRepositoryInterface interface {
	Create(provider *model.Provider) error
	GetAll() ([]model.Provider, error)
	// Update(id string, updates map[string]any) error
	// Delete(id string) error
}
