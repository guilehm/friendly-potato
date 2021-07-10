package main

import (
	"fmt"
	"goapi/handlers"
	"goapi/middlewares"
	"goapi/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("hello from goapi")

	r := mux.NewRouter()
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(r)
	r.StrictSlash(true).HandleFunc("/games/", middlewares.Authentication(handlers.GameList))
	r.StrictSlash(true).HandleFunc("/games/{id}/", handlers.GameDetail)
	r.StrictSlash(true).HandleFunc("/users/", handlers.SignUp).Methods("POST")
	r.StrictSlash(true).HandleFunc("/users/login/", handlers.Login).Methods("POST")
	http.ListenAndServe(":"+os.Getenv("PORT"), utils.SetHeaders(utils.LogRequest(handler)))
}
