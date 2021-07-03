package handlers

import (
	"encoding/json"
	"fmt"
	"goapi/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Games(w http.ResponseWriter, r *http.Request) {
	rawg := services.RawgService()
	resp, err := rawg.SearchGame(r.URL.Query())

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GamesDetail(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
