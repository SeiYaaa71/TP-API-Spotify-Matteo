package controller

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
)

type SearchAPI struct {
    Tracks struct {
        Items []struct {
            Name string `json:"name"`
            Album struct {
                Name   string `json:"name"`
                Images []struct{ URL string `json:"url"` } `json:"images"`
                Release string `json:"release_date"`
            } `json:"album"`
            ExternalURLs struct {
                Spotify string `json:"spotify"`
            } `json:"external_urls"`
            Artists []struct {
                Name string `json:"name"`
            } `json:"artists"`
        } `json:"items"`
    } `json:"tracks"`
}

type TrackData struct {
    Title       string
    AlbumName   string
    AlbumCover  string
    ReleaseDate string
    Artist      string
    SpotifyURL  string
}

func LaylowTrack(w http.ResponseWriter, r *http.Request) {

    url := "https://api.spotify.com/v1/search?q=laylow%20maladresse&type=track&market=FR&limit=1"

    data, err := SpotifyGET(url)
    if err != nil {
        http.Error(w, "Erreur API Spotify", 500)
        return
    }

    var api SearchAPI
    json.Unmarshal(data, &api)

    if len(api.Tracks.Items) == 0 {
        http.Error(w, "Aucun titre trouvé", 404)
        return
    }

    item := api.Tracks.Items[0]

    cover := ""
    if len(item.Album.Images) > 0 {
        cover = item.Album.Images[0].URL
    }

    track := TrackData{
        Title:       item.Name,
        AlbumName:   item.Album.Name,
        AlbumCover:  cover,
        ReleaseDate: item.Album.Release,
        Artist:      item.Artists[0].Name,
        SpotifyURL:  item.ExternalURLs.Spotify,
    }

    log.Println("Track récupérée :", track.Title)

    tmpl, _ := template.ParseFiles("template/laylow.html")
    tmpl.Execute(w, track)
}
