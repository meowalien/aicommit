package ai

import "fmt"

const baseSystemPrompt = `You are an expert at writing git commit messages following the Conventional Commits specification.

Given a git diff, generate a concise and descriptive commit message.

Rules:
1. Use the format: <type>(<scope>): <description>
2. Types: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert
3. Scope is optional but recommended when changes are focused on a specific area
4. Description should be imperative mood, lowercase, no period at the end
5. Keep the first line under 72 characters
6. If the changes are complex, add a blank line followed by a more detailed body explaining what and why

Respond with ONLY the commit message, no explanations, no markdown formatting, no quotes around the message.`

func BuildSystemPrompt(language string) string {
	if language == "" || language == "en" {
		return baseSystemPrompt
	}

	return fmt.Sprintf(`%s

IMPORTANT: Write the commit message in %s language. The type (feat, fix, etc.) should remain in English, but the description and body should be written in %s.`, baseSystemPrompt, language, language)
}

func BuildUserPrompt(diff string) string {
	// Truncate diff if too long to avoid token limits
	const maxDiffLength = 10000
	if len(diff) > maxDiffLength {
		diff = diff[:maxDiffLength] + "\n\n... (diff truncated)"
	}

	return fmt.Sprintf("Generate a commit message for the following git diff:\n\n%s", diff)
}
