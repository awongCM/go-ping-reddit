package main

import "fmt"

type demoPost struct {
	Author string
	Title  string
}

// sample posts shaped like /r/golang listings for offline demo runs.
var demoPosts = map[string][]demoPost{
	"golang": {
		{Author: "gopher_dev", Title: "Go 1.23 release candidate is out"},
		{Author: "std_lib_fan", Title: "Why I still reach for channels over mutexes"},
		{Author: "backend_owl", Title: "Benchmarking HTTP handlers with testing.B"},
		{Author: "module_maven", Title: "A practical guide to Go workspaces in monorepos"},
		{Author: "concurrency_cat", Title: "context.Context patterns that saved my production service"},
	},
	"programming": {
		{Author: "lang_nerd", Title: "What language would you pick for a new CLI tool in 2026?"},
		{Author: "refactor_guru", Title: "Small refactors that pay off immediately"},
		{Author: "api_architect", Title: "REST vs GraphQL for internal tools"},
		{Author: "test_driver", Title: "Integration tests without a dedicated staging stack"},
		{Author: "debug_duck", Title: "The bug was a timezone. It is always a timezone."},
	},
}

func runDemo(subreddit string, limit int) error {
	if err := validateLimit(limit); err != nil {
		return err
	}
	if err := validateSubreddit(subreddit); err != nil {
		return err
	}

	sourceSubreddit := subreddit
	posts, ok := demoPosts[subreddit]
	if !ok {
		sourceSubreddit = "golang"
		posts = demoPosts["golang"]
		fmt.Printf("demo: no sample data for /r/%s, using /r/%s instead\n", subreddit, sourceSubreddit)
	}

	limit = capLimit(limit, len(posts))
	fmt.Printf("demo mode — sample posts from /r/%s (no Reddit API call)\n\n", sourceSubreddit)
	for _, post := range posts[:limit] {
		fmt.Printf("[%s] %s\n", post.Author, post.Title)
	}

	return nil
}
