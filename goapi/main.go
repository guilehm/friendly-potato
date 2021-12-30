package main

import (
	"fmt"
	"goapi/db"
	"goapi/handlers"
	"goapi/middlewares"
	"goapi/ws"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("hello from goapi")
	db.CreateIndexes()

	r := mux.NewRouter()
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(r)
	r.StrictSlash(true).HandleFunc("/games/", handlers.GameList)
	r.StrictSlash(true).HandleFunc("/games/{slug}/", handlers.GameDetail)
	r.StrictSlash(true).HandleFunc("/users/", handlers.SignUp).Methods("POST")
	r.StrictSlash(true).HandleFunc("/users/login/", handlers.Login).Methods("POST")
	r.StrictSlash(true).HandleFunc("/users/refresh/", handlers.RefreshToken).Methods("POST")
	r.StrictSlash(true).HandleFunc("/users/validate/", handlers.ValidateToken).Methods("POST")

	r.HandleFunc("/socket/", ws.SocketHandler)

	_ = http.ListenAndServe(":"+os.Getenv("PORT"), middlewares.SetHeaders(middlewares.LogRequest(handler)))
}
