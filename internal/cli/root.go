package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/jacky_li/aicommit/internal/ai"
	"github.com/jacky_li/aicommit/internal/config"
	"github.com/jacky_li/aicommit/internal/git"
)

var (
	dryRun  bool
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "aicommit",
	Short: "Generate AI-powered git commit messages",
	Long: `aicommit uses AI to generate Conventional Commits messages from your staged changes.

Usage:
  aicommit              Generate commit message and commit staged changes
  aicommit set KEY=VAL  Configure API keys and settings

Examples:
  # Set up Anthropic
  aicommit set anthropic_key=sk-ant-xxx
  aicommit set provider=anthropic

  # Set up OpenAI
  aicommit set openai_key=sk-xxx
  aicommit set provider=openai

  # Generate and commit
  git add .
  aicommit`,
	RunE: runCommit,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show generated message without committing")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
}

func runCommit(cmd *cobra.Command, args []string) error {
	// 1. Check if in git repository
	if !git.IsGitRepo() {
		return fmt.Errorf("not a git repository")
	}

	// 2. Check for staged changes
	hasChanges, err := git.HasStagedChanges()
	if err != nil {
		return fmt.Errorf("failed to check staged changes: %w", err)
	}
	if !hasChanges {
		return fmt.Errorf("no staged changes found. Stage changes with 'git add' first")
	}

	// 3. Load configuration
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// 4. Validate configuration
	if err := cfg.Validate(); err != nil {
		return err
	}

	// 5. Get staged diff
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("failed to get diff: %w", err)
	}

	if verbose {
		files, _ := git.GetStagedFiles()
		fmt.Printf("Staged files: %v\n", files)
	}

	// 6. Create AI provider
	provider, err := ai.NewProvider(cfg)
	if err != nil {
		return fmt.Errorf("failed to create AI provider: %w", err)
	}

	// 7. Generate commit message
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Printf("Generating commit message using %s...\n", provider.Name())
	message, err := provider.GenerateCommitMessage(ctx, diff)
	if err != nil {
		return fmt.Errorf("failed to generate commit message: %w", err)
	}

	// 8. Show message
	fmt.Printf("\nGenerated commit message:\n%s\n", message)

	if dryRun {
		fmt.Println("\n(dry-run mode - not committing)")
		return nil
	}

	// 9. Commit
	if err := git.Commit(message); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	fmt.Println("\nCommit successful!")
	return nil
}
