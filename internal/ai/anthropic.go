package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AnthropicProvider struct {
	client anthropic.Client
	model  string
}

func NewAnthropicProvider(apiKey, model string) (*AnthropicProvider, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("Anthropic API key not configured")
	}
	if model == "" {
		model = "claude-sonnet-4-20250514"
	}
	return &AnthropicProvider{
		client: anthropic.NewClient(option.WithAPIKey(apiKey)),
		model:  model,
	}, nil
}

func (p *AnthropicProvider) Name() string {
	return "anthropic"
}

func (p *AnthropicProvider) GenerateCommitMessage(ctx context.Context, diff string) (string, error) {
	message, err := p.client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:     anthropic.Model(p.model),
		MaxTokens: 500,
		System: []anthropic.TextBlockParam{
			{Text: SystemPrompt},
		},
		Messages: []anthropic.MessageParam{
			{
				Role: anthropic.MessageParamRoleUser,
				Content: []anthropic.ContentBlockParamUnion{
					anthropic.NewTextBlock(BuildUserPrompt(diff)),
				},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("Anthropic API error: %w", err)
	}

	if len(message.Content) == 0 {
		return "", fmt.Errorf("no response from Anthropic")
	}

	// Extract text from response
	for _, block := range message.Content {
		if block.Type == "text" {
			return strings.TrimSpace(block.Text), nil
		}
	}

	return "", fmt.Errorf("no text response from Anthropic")
}
