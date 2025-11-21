package controller

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
)

type AlbumAPI struct {
    Items []struct {
        Name        string `json:"name"`
        ReleaseDate string `json:"release_date"`
        TotalTracks int    `json:"total_tracks"`
        Images      []struct {
            URL    string `json:"url"`
            Height int    `json:"height"`
            Width  int    `json:"width"`
        } `json:"images"`
    } `json:"items"`
}


type AlbumData struct {
    Name        string
    ReleaseDate string
    TrackCount  int
    Image       string
}

func DamsoAlbums(w http.ResponseWriter, r *http.Request) {

    url := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"

    data, err := SpotifyGET(url)
    if err != nil {
        http.Error(w, "Erreur Spotify", 500)
        return
    }

    // ★ DEBUG ICI — data est bien défini à cet endroit
    log.Println("JSON Spotify reçu :", string(data))

    var apiRes AlbumAPI
    if err := json.Unmarshal(data, &apiRes); err != nil {
        log.Println("ERREUR JSON:", err)
        http.Error(w, "Erreur JSON Spotify", 500)
        return
    }

    log.Println("Albums reçus :", len(apiRes.Items)) // DEBUG

    var albums []AlbumData

    for _, item := range apiRes.Items {

        img := ""
        if len(item.Images) > 0 {
            img = item.Images[0].URL
        }

        albums = append(albums, AlbumData{
            Name:        item.Name,
            ReleaseDate: item.ReleaseDate,
            TrackCount:  item.TotalTracks,
            Image:       img,
        })
    }

    tmpl := template.Must(template.ParseFiles("template/damso.html"))
    tmpl.Execute(w, albums)
}
