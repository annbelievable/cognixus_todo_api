package controllers

import (
	"cognixus_todo_api/inits"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type UserInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func HandleGoogleLogin(c *gin.Context) {
	url := inits.GoogleOauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "state" {
		c.String(http.StatusBadRequest, "Invalid state parameter")
		return
	}

	code := c.Query("code")
	token, err := inits.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	userInfo, err := getUserInfo(token)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info")
		return
	}

	session := sessions.Default(c)
	session.Set("userEmail", userInfo.Email)
	session.Save()

	c.String(http.StatusOK, "User name: %s, user email: %s", userInfo.Name, userInfo.Email)
}

func getUserInfo(token *oauth2.Token) (*UserInfo, error) {
	client := inits.GoogleOauthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
