package services

import (
	"net/http"
	"os"

	"github.com/josephrockqz/wemusic-golang/pkg/utils"
	"github.com/labstack/echo/v4"
)

func SpotifyLogin(context echo.Context) error {
	redirectLocation := constructRedirectLocation()

	return context.Redirect(http.StatusPermanentRedirect, redirectLocation)
}

func constructRedirectLocation() string {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	state := utils.GenerateRandomString(16)

	url := "https://accounts.spotify.com/authorize"
	url += "?client_id=" + clientId
	url += "&response_type=code"
	url += "&redirect_uri=http://localhost:8080/spotify-user-authorization-callback"
	url += "&scope=user-read-private%20user-read-email"
	url += "&state=" + state

	return url
}
