package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AnthropicConfig struct {
	APIKey string `mapstructure:"api_key"`
	Model  string `mapstructure:"model"`
}

type OpenAIConfig struct {
	APIKey string `mapstructure:"api_key"`
	Model  string `mapstructure:"model"`
}

type Config struct {
	Provider  string          `mapstructure:"provider"`
	Language  string          `mapstructure:"language"`
	Anthropic AnthropicConfig `mapstructure:"anthropic"`
	OpenAI    OpenAIConfig    `mapstructure:"openai"`
}

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".aicommit"), nil
}

func getConfigPath() (string, error) {
	dir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

func Default() *Config {
	return &Config{
		Provider: "anthropic",
		Language: "en",
		Anthropic: AnthropicConfig{
			Model: "claude-sonnet-4-20250514",
		},
		OpenAI: OpenAIConfig{
			Model: "gpt-4o",
		},
	}
}

func Load() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	// Set defaults
	viper.SetDefault("provider", "anthropic")
	viper.SetDefault("language", "en")
	viper.SetDefault("anthropic.model", "claude-sonnet-4-20250514")
	viper.SetDefault("openai.model", "gpt-4o")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, return default config
			return Default(), nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) Save() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	viper.Set("provider", c.Provider)
	viper.Set("language", c.Language)
	viper.Set("anthropic.api_key", c.Anthropic.APIKey)
	viper.Set("anthropic.model", c.Anthropic.Model)
	viper.Set("openai.api_key", c.OpenAI.APIKey)
	viper.Set("openai.model", c.OpenAI.Model)

	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	// Set file permissions to owner only
	if err := os.Chmod(configPath, 0600); err != nil {
		return fmt.Errorf("failed to set config permissions: %w", err)
	}

	return nil
}

func (c *Config) Set(key, value string) error {
	switch key {
	case "provider":
		if value != "openai" && value != "anthropic" {
			return fmt.Errorf("invalid provider: %s (must be 'openai' or 'anthropic')", value)
		}
		c.Provider = value
	case "language", "lang":
		c.Language = value
	case "anthropic_key":
		c.Anthropic.APIKey = value
	case "anthropic_model":
		c.Anthropic.Model = value
	case "openai_key":
		c.OpenAI.APIKey = value
	case "openai_model":
		c.OpenAI.Model = value
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}
	return nil
}

func (c *Config) Validate() error {
	if c.Provider == "" {
		return fmt.Errorf("provider not configured")
	}

	switch c.Provider {
	case "anthropic":
		if c.Anthropic.APIKey == "" {
			return fmt.Errorf("Anthropic API key not configured. Run: aicommit set anthropic_key=YOUR_KEY")
		}
	case "openai":
		if c.OpenAI.APIKey == "" {
			return fmt.Errorf("OpenAI API key not configured. Run: aicommit set openai_key=YOUR_KEY")
		}
	default:
		return fmt.Errorf("invalid provider: %s", c.Provider)
	}

	return nil
}
