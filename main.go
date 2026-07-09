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
