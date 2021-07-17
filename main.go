package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := oauth1.NewConfig(os.Getenv("KEY"), os.Getenv("SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	log.Print(os.Getenv("KEY"))
	log.Print(os.Getenv("SECRET"))
	log.Print(os.Getenv("ACCESS_TOKEN"))
	log.Print(os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("just setting up my twttr bot", nil)

	if err == nil {
		fmt.Printf("tweet: %-v", tweet)
		fmt.Printf("response: %-v", resp)
	} else {
		fmt.Printf("error: %-v", err)
	}
}
