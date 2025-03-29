package service

import "github.com/wengchaoxi/one-suno-api/internal/provider"

type ServiceOptions struct {
	ProviderManager provider.ProviderManager
	ApiKey          string
}

type Service struct {
	opts *ServiceOptions
}

func New(opts *ServiceOptions) *Service {
	return &Service{
		opts: opts,
	}
}
