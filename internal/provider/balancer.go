package provider

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randSrc = rand.NewSource(time.Now().UnixNano())
	randMu  sync.Mutex
)

type ProviderBalancer interface {
	Select(providers []ProviderMeta) *ProviderMeta
}

type WeightedRoundRobinBalancer struct{}

func NewWeightedRoundRobinBalancer() *WeightedRoundRobinBalancer {
	return &WeightedRoundRobinBalancer{}
}

func (w *WeightedRoundRobinBalancer) Select(providers []ProviderMeta) *ProviderMeta {
	total := 0
	for _, p := range providers {
		total += p.Weight
	}

	randMu.Lock()
	r := rand.New(randSrc).Intn(total)
	randMu.Unlock()

	sum := 0
	for _, p := range providers {
		sum += p.Weight
		if sum > r {
			return &p
		}
	}
	return nil
}
