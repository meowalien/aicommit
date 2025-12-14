# aicommit

AI-powered git commit message generator using OpenAI or Anthropic APIs.

## Features

- Automatically generates [Conventional Commits](https://www.conventionalcommits.org/) formatted commit messages
- Supports both OpenAI and Anthropic AI providers
- Simple CLI interface
- Secure configuration storage

## Installation

```bash
# Clone and build
git clone https://github.com/jacky_li/aicommit.git
cd aicommit
go build -o aicommit ./cmd/aicommit/

# Move to PATH (optional)
sudo mv aicommit /usr/local/bin/
```

## Configuration

Set your API keys and choose a provider:

```bash
# Using Anthropic (default)
aicommit set anthropic_key=sk-ant-xxxxx
aicommit set provider=anthropic

# Using OpenAI
aicommit set openai_key=sk-xxxxx
aicommit set provider=openai

# Optional: specify a different model
aicommit set anthropic_model=claude-sonnet-4-20250514
aicommit set openai_model=gpt-4o
```

Configuration is stored in `~/.aicommit/config.yaml`.

## Usage

```bash
# Stage your changes
git add .

# Generate commit message and commit
aicommit

# Preview generated message without committing
aicommit --dry-run

# Verbose output
aicommit --verbose
```

## Available Commands

```
aicommit              Generate commit message and commit staged changes
aicommit set KEY=VAL  Configure API keys and settings
aicommit --help       Show help
```

## Configuration Keys

| Key | Description | Default |
|-----|-------------|---------|
| `provider` | AI provider (`openai` or `anthropic`) | `anthropic` |
| `anthropic_key` | Anthropic API key | - |
| `anthropic_model` | Anthropic model | `claude-sonnet-4-20250514` |
| `openai_key` | OpenAI API key | - |
| `openai_model` | OpenAI model | `gpt-4o` |

## Example Output

```
$ git add .
$ aicommit
Generating commit message using anthropic...

Generated commit message:
feat(auth): add user login validation

Commit successful!
```

## License

MIT
