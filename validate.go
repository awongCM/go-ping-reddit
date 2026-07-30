package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func validateLimit(limit int) error {
	if limit < 1 {
		return fmt.Errorf("limit must be at least 1, got %d", limit)
	}
	return nil
}

func capLimit(limit, max int) int {
	if limit > max {
		return max
	}
	return limit
}

func validateSubreddit(name string) error {
	if name == "" {
		return fmt.Errorf("subreddit name cannot be empty")
	}

	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			continue
		}
		return fmt.Errorf("invalid subreddit name %q: use only letters, numbers, and underscores", name)
	}

	return nil
}

func validateAgentFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read agent file: %w", err)
	}

	placeholders := []string{
		"<client id",
		"<client secret>",
		"<reddit username>",
		"<reddit password>",
	}
	for _, placeholder := range placeholders {
		if strings.Contains(string(content), placeholder) {
			return fmt.Errorf(
				"agent file %q still contains placeholder values; edit it with your Reddit credentials (see README)",
				path,
			)
		}
	}

	return nil
}
