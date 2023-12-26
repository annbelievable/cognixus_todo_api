package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const googleTokenInfoURL = "https://www.googleapis.com/oauth2/v3/tokeninfo"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userEmail := session.Get("userEmail")
		if userEmail == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login-required/")
			c.Abort()
			return
		}

		githubToken := session.Get("githubToken")
		googleToken := session.Get("googleToken")

		if githubToken == nil && googleToken == nil {
			c.Redirect(http.StatusUnauthorized, "/login-required/")
			c.Abort()
			return
		}

		if githubToken != nil && !verifyGitHubToken(githubToken.(string)) {
			c.Redirect(http.StatusUnauthorized, "/login-required/")
			c.Abort()
			return
		}

		if googleToken != nil && !verifyGoogleToken(googleToken.(string)) {
			c.Redirect(http.StatusUnauthorized, "/login-required/")
			c.Abort()
			return
		}

		c.Next()
	}
}

func verifyGitHubToken(token string) bool {
	client := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
	_, _, err := client.Users.Get(context.Background(), "")
	return err == nil
}

func verifyGoogleToken(token string) bool {
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", googleTokenInfoURL, token))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var tokenInfo map[string]interface{}
	if err := json.Unmarshal(body, &tokenInfo); err != nil {
		return false
	}

	if resp.StatusCode == http.StatusOK && tokenInfo["error"] == nil {
		return true
	}

	return false
}
