package controllers

import (
	"cognixus_todo_api/inits"
	"context"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func HandleGithubLogin(c *gin.Context) {
	url := inits.GithubOauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGithubCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "state" {
		c.String(http.StatusBadRequest, "Invalid state parameter")
		return
	}

	code := c.Query("code")
	token, err := inits.GithubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token.AccessToken})))
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info")
		return
	}

	emails, _, err := client.Users.ListEmails(context.Background(), nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var primaryEmail string
	for _, email := range emails {
		if email.GetPrimary() {
			primaryEmail = email.GetEmail()
			break
		}
	}

	session := sessions.Default(c)
	session.Set("userEmail", primaryEmail)
	session.Set("githubToken", token.AccessToken)
	session.Save()

	c.String(http.StatusOK, "User name: %s, user email: %s", user.GetLogin(), primaryEmail)
}
