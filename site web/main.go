package main

import (
    "log"
    "net/http"

    "TpSpotify/controller"
    "TpSpotify/router"
)

func main() {
    controller.InitSpotify()

    r := router.New()

    log.Println("Serveur lancé sur : http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
