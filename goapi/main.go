package main

import (
	"fmt"
	"goapi/db"
	"goapi/handlers"
	"goapi/middlewares"
	"goapi/models"
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

	r.HandleFunc("/ws/lumber/", handlers.SocketLumberHandler)

	hub := models.Hub{
		Broadcast:  make(chan bool),
		Register:   make(chan *models.Client),
		Unregister: make(chan *models.Client),
		Clients:    make(map[*models.Client]bool),
	}
	go hub.Start()
	r.HandleFunc("/ws/rpg/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RPGHandler(&hub, w, r)
	})

	_ = http.ListenAndServe(":"+os.Getenv("PORT"), middlewares.SetHeaders(middlewares.LogRequest(handler)))
}
