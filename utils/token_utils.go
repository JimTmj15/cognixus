package utils

import (
	"context"

	"github.com/google/go-github/v55/github"
)

func ValidateToken(c context.Context, token string) (int64, error) {
	client := github.NewClient(nil).WithAuthToken(token)
	user, _, err := client.Users.Get(c, "")
	if err != nil {
		return 0, err
	}

	return user.GetID(), nil
}
