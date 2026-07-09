package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/turnage/graw/reddit"
)

func main() {
	agentFile := flag.String("agent", "reddit-account.agent", "path to Reddit credentials file")
	subreddit := flag.String("sub", "golang", "subreddit to read (without /r/)")
	limit := flag.Int("limit", 5, "number of posts to print")
	demo := flag.Bool("demo", false, "run with sample data (no Reddit credentials required)")
	flag.Parse()

	if *demo {
		runDemo(*subreddit, *limit)
		return
	}

	bot, err := reddit.NewBotFromAgentFile(*agentFile, 0)
	if err != nil {
		log.Fatalf("create bot: %v", err)
	}

	path := fmt.Sprintf("/r/%s", *subreddit)
	harvest, err := bot.Listing(path, "")
	if err != nil {
		log.Fatalf("fetch %s: %v", path, err)
	}

	if len(harvest.Posts) == 0 {
		fmt.Println("no posts found")
		return
	}

	n := *limit
	if n > len(harvest.Posts) {
		n = len(harvest.Posts)
	}

	for _, post := range harvest.Posts[:n] {
		fmt.Printf("[%s] %s\n", post.Author, post.Title)
	}
}

func runDemo(subreddit string, limit int) {
	fmt.Printf("go-ping-reddit demo — /r/%s (sample data, no API call)\n\n", subreddit)

	posts := demoPosts(subreddit)
	n := limit
	if n > len(posts) {
		n = len(posts)
	}

	for _, post := range posts[:n] {
		fmt.Printf("[%s] %s\n", post.author, post.title)
	}
}

type demoPost struct {
	author string
	title  string
}

func demoPosts(subreddit string) []demoPost {
	switch subreddit {
	case "golang":
		return []demoPost{
			{"gopher_fan", "What's your favorite Go 1.22 feature?"},
			{"stdlib_lover", "Benchmarking sync.Pool vs channel pools"},
			{"modules_user", "go mod tidy keeps removing my test deps"},
			{"concurrency_nerd", "errgroup vs waitgroup in production"},
			{"newbie_dev", "Best resources for learning Go in 2026?"},
		}
	default:
		return []demoPost{
			{"demo_user", fmt.Sprintf("Sample post 1 from /r/%s", subreddit)},
			{"demo_user", fmt.Sprintf("Sample post 2 from /r/%s", subreddit)},
			{"demo_user", fmt.Sprintf("Sample post 3 from /r/%s", subreddit)},
		}
	}
}
