package service

import (
	"github.com/wengchaoxi/one-suno-api/internal/provider"
	"github.com/wengchaoxi/one-suno-api/internal/repository"
)

type ServiceOptions struct {
	ProviderManager  provider.ProviderManager
	ApiKey           string
	UserRepository   repository.UserRepositoryInterface
	ApiKeyRepository repository.ApiKeyRepositoryInterface
}

type Service struct {
	opts *ServiceOptions
}

func New(opts *ServiceOptions) *Service {
	return &Service{
		opts: opts,
	}
}
