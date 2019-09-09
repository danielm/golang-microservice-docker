package rest

import (
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

type ApiService struct {
    router http.Handler
}

func (a *ApiService) salute(w http.ResponseWriter, r *http.Request) {
    log.Print("Got a hit at 'salute' endpoint")

    w.Header().Set("Content-Type", "application/json")

    message := map[string]string{
        "status": "OK",
        "message": "Hello World",
    }

    json.NewEncoder(w).Encode(message)
}

func (a *ApiService) GetRouter() http.Handler {
    return a.router
}

func New() ApiService {
    api := ApiService{}

    router := mux.NewRouter()
    router.HandleFunc("/", api.salute).Methods(http.MethodGet)

    api.router = router

    return api
}
