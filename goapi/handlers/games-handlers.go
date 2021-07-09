package handlers

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	"goapi/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GameList(w http.ResponseWriter, r *http.Request) {
	rawg := services.RawgService()
	resp, err := rawg.SearchGame(r.URL.Query())

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	db.OpenCollection("game-list")

	w.Write(jsonResponse)
}

func GameDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	rawg := services.RawgService()
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: invalid ID")
	}

	resp, err := rawg.GetGameDetail(intId)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	w.Write(jsonResponse)
}
