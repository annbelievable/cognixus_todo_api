package inits

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config

func SetGoogleOauthConfig() {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	GoogleOauthConfig = &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		RedirectURL:  "http://localhost:3000/callback/google/",
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
