package router

import (
    "net/http"

    "TpSpotify/controller"
)

func New() *http.ServeMux {
    mux := http.NewServeMux()

    // ROUTE ACCUEIL
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "template/index.html")
    })

    // ROUTES
    mux.HandleFunc("/album/damso", controller.DamsoAlbums)
    mux.HandleFunc("/track/laylow", controller.LaylowTrack)

    // SERVIR LE CSS
    fs := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))

    return mux
}
