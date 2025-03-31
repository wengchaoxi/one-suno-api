package provider

import "github.com/wengchaoxi/one-suno-api/internal/dto"

type Provider interface {
	CreateAudio(*dto.CreateAudioRequest) (*dto.CreateAudioResponse, error)
}

type ProviderMeta struct {
	Id       string
	Weight   int
	Provider Provider
}

type ProviderManager struct {
	providers []ProviderMeta
	Balancer  ProviderBalancer
}

func NewProviderManager(balancer ProviderBalancer) *ProviderManager {
	return &ProviderManager{
		Balancer: balancer,
	}
}

func (p *ProviderManager) Register(provider ...ProviderMeta) {
	p.providers = append(p.providers, provider...)
}

func (p *ProviderManager) Select() *ProviderMeta {
	return p.Balancer.Select(p.providers)
}
