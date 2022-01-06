package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func PostTweet(auth *TwitterAuth) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var config *oauth1.Config
	var token *oauth1.Token
	if auth == nil {
		config = oauth1.NewConfig(os.Getenv("KEY"), os.Getenv("SECRET"))
		token = oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	} else {
		config = auth.Config
		token = auth.Token
	}
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("Since the 1940s ... fastest year: x (thanks artist), slowest year: ... trending up/down now", nil)

	if err == nil {
		fmt.Printf("tweet: %-v", tweet)
		fmt.Printf("response: %-v", resp)
	} else {
		fmt.Printf("error: %-v", err)
	}
}

type TwitterAuth struct {
	Config *oauth1.Config
	Token  *oauth1.Token
}
