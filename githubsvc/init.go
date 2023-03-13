package githubsvc

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var githubClient *github.Client

func initDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func initGithubClient() {
	log.Println("initializing github client")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient = github.NewClient(tc)
}

func init() {
	initDotEnv()
	initGithubClient()
}
