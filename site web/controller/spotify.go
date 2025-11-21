package controller

import (
    "encoding/json"
    "io"
    "net/http"
    "os"
    "strings"
	"fmt"
)

var spotifyToken string

func InitSpotify() {
    clientID := os.Getenv("SPOTIFY_CLIENT_ID")
    clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

    reqBody := strings.NewReader("grant_type=client_credentials")
    req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", reqBody)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.SetBasicAuth(clientID, clientSecret)

    resp, _ := http.DefaultClient.Do(req)
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var t struct {
        AccessToken string `json:"access_token"`
    }
    json.Unmarshal(body, &t)

    spotifyToken = t.AccessToken
}

func SpotifyGET(url string) ([]byte, error) {
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+spotifyToken)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
