package main

import (
	"fmt"
	"goapi/handlers"
	"goapi/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello from gorawg")

	r := mux.NewRouter()
	r.StrictSlash(true).HandleFunc("/games/", handlers.Games)
	r.StrictSlash(true).HandleFunc("/games/{id}/", handlers.GamesDetail)
	r.StrictSlash(true).HandleFunc("/users/", handlers.SignUp)
	http.ListenAndServe(":"+os.Getenv("PORT"), utils.LogRequest(r))
}
