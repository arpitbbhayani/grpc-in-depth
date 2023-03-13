package githubsvc

import (
	"context"

	"github.com/google/go-github/github"
)

func Me() *github.User {
	ctx := context.Background()
	user, _, err := githubClient.Users.Get(ctx, "")
	if err != nil {
		panic(err)
	}
	return user
}
