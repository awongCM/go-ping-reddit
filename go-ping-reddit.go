// TODO

package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/turnage/graw/reddit"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {

	requestToken()

	bot, err := reddit.NewBotFromAgentFile("reddit-account.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}

//see txt file - if any
func requestToken() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://www.reddit.com/api/v1/access_token"},
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit the url for the url auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}
