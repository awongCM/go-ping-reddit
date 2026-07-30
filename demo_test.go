package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureDemoOutput(t *testing.T, subreddit string, limit int) (string, error) {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	runErr := runDemo(subreddit, limit)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	r.Close()

	return buf.String(), runErr
}

func TestRunDemoKnownSubreddit(t *testing.T) {
	t.Parallel()

	out, err := captureDemoOutput(t, "golang", 2)
	if err != nil {
		t.Fatalf("runDemo() error = %v", err)
	}
	if !strings.Contains(out, "demo mode — sample posts from /r/golang") {
		t.Fatalf("output missing golang banner: %q", out)
	}
	if !strings.Contains(out, "[gopher_dev]") {
		t.Fatalf("output missing sample post: %q", out)
	}
}

func TestRunDemoUnknownSubredditFallback(t *testing.T) {
	t.Parallel()

	out, err := captureDemoOutput(t, "rust", 2)
	if err != nil {
		t.Fatalf("runDemo() error = %v", err)
	}
	if !strings.Contains(out, "no sample data for /r/rust, using /r/golang instead") {
		t.Fatalf("output missing fallback notice: %q", out)
	}
	if !strings.Contains(out, "demo mode — sample posts from /r/golang") {
		t.Fatalf("output should show golang as source, got: %q", out)
	}
	if strings.Contains(out, "demo mode — sample posts from /r/rust") {
		t.Fatalf("output should not claim rust as source: %q", out)
	}
}

func TestRunDemoInvalidLimit(t *testing.T) {
	t.Parallel()

	_, err := captureDemoOutput(t, "golang", -1)
	if err == nil {
		t.Fatal("runDemo(-1) expected error")
	}
}

func TestRunDemoInvalidSubreddit(t *testing.T) {
	t.Parallel()

	_, err := captureDemoOutput(t, "foo/bar", 2)
	if err == nil {
		t.Fatal("runDemo(foo/bar) expected error")
	}
}
