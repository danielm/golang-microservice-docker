package main

import (
    "os"
    "log"
    "time"
    "net/http"
    "./rest"
)

func main() {
    api := rest.New()

    server := http.Server{
        Handler:        api.GetRouter(),
        Addr:           ":8080",
        WriteTimeout:   time.Second,
        ReadTimeout:    time.Second,
    }

    log.Print("Starting: ", os.Getenv("PROJECT_NAME"))
    log.Print("Environment is: ", os.Getenv("ENV"))
    log.Print("Exposed is: ", os.Getenv("HTTP_PORT"))

    log.Fatal(server.ListenAndServe())
}
