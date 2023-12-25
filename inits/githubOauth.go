package inits

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOauthConfig *oauth2.Config

func SetGithubOauthConfig() {
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	GithubOauthConfig = &oauth2.Config{
		ClientID:     githubClientID,
		ClientSecret: githubClientSecret,
		RedirectURL:  "http://localhost:3000/callback/github/",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}
