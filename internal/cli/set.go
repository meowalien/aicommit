package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jacky_li/aicommit/internal/config"
)

var setCmd = &cobra.Command{
	Use:   "set [key=value ...]",
	Short: "Set configuration values",
	Long: `Set configuration values for aicommit.

Available keys:
  provider          - AI provider to use (openai or anthropic)
  anthropic_key     - Anthropic API key
  anthropic_model   - Anthropic model name (default: claude-sonnet-4-20250514)
  openai_key        - OpenAI API key
  openai_model      - OpenAI model name (default: gpt-4o)

Examples:
  aicommit set provider=anthropic
  aicommit set anthropic_key=sk-ant-xxx anthropic_model=claude-sonnet-4-20250514
  aicommit set openai_key=sk-xxx openai_model=gpt-4o`,
	RunE: runSet,
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func runSet(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no key=value pairs provided\n\nUsage: aicommit set key=value [key=value ...]")
	}

	cfg, err := config.Load()
	if err != nil {
		// If config doesn't exist, create new
		cfg = config.Default()
	}

	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid format: %s (expected key=value)", arg)
		}

		key, value := parts[0], parts[1]
		if err := cfg.Set(key, value); err != nil {
			return fmt.Errorf("failed to set %s: %w", key, err)
		}
		fmt.Printf("Set %s\n", key)
	}

	if err := cfg.Save(); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Println("Configuration saved to ~/.aicommit/config.yaml")
	return nil
}
