package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		limit   int
		wantErr bool
	}{
		{name: "positive", limit: 5},
		{name: "one", limit: 1},
		{name: "zero", limit: 0, wantErr: true},
		{name: "negative", limit: -1, wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validateLimit(tt.limit)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateLimit(%d) error = %v, wantErr %v", tt.limit, err, tt.wantErr)
			}
		})
	}
}

func TestCapLimit(t *testing.T) {
	t.Parallel()

	if got := capLimit(10, 3); got != 3 {
		t.Fatalf("capLimit(10, 3) = %d, want 3", got)
	}
	if got := capLimit(2, 5); got != 2 {
		t.Fatalf("capLimit(2, 5) = %d, want 2", got)
	}
}

func TestValidateSubreddit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		sub     string
		wantErr bool
	}{
		{name: "valid", sub: "golang"},
		{name: "underscore", sub: "learn_python"},
		{name: "empty", sub: "", wantErr: true},
		{name: "slash", sub: "foo/bar", wantErr: true},
		{name: "spaces", sub: "foo bar", wantErr: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validateSubreddit(tt.sub)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateSubreddit(%q) error = %v, wantErr %v", tt.sub, err, tt.wantErr)
			}
		})
	}
}

func TestValidateAgentFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()

	valid := filepath.Join(dir, "valid.agent")
	if err := os.WriteFile(valid, []byte(`user_agent: "linux:test:0.1.0 (by /u/demo)"
client_id: "abc123"
client_secret: "secret"
username: "demo"
password: "pass"
`), 0o600); err != nil {
		t.Fatal(err)
	}

	placeholder := filepath.Join(dir, "placeholder.agent")
	if err := os.WriteFile(placeholder, []byte(`user_agent: "linux:test:0.1.0 (by /u/<reddit username>)"
client_id: "<client id from reddit.com/prefs/apps>"
client_secret: "<client secret>"
username: "<reddit username>"
password: "<reddit password>"
`), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := validateAgentFile(valid); err != nil {
		t.Fatalf("validateAgentFile(valid) = %v", err)
	}
	if err := validateAgentFile(placeholder); err == nil {
		t.Fatal("validateAgentFile(placeholder) expected error")
	}
	if err := validateAgentFile(filepath.Join(dir, "missing.agent")); err == nil {
		t.Fatal("validateAgentFile(missing) expected error")
	}
}
