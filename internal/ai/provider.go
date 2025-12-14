package ai

import (
	"context"
	"fmt"

	"github.com/jacky_li/aicommit/internal/config"
)

type Provider interface {
	GenerateCommitMessage(ctx context.Context, diff string) (string, error)
	Name() string
}

func NewProvider(cfg *config.Config) (Provider, error) {
	switch cfg.Provider {
	case "openai":
		return NewOpenAIProvider(cfg.OpenAI.APIKey, cfg.OpenAI.Model)
	case "anthropic":
		return NewAnthropicProvider(cfg.Anthropic.APIKey, cfg.Anthropic.Model)
	default:
		return nil, fmt.Errorf("unknown provider: %s", cfg.Provider)
	}
}
