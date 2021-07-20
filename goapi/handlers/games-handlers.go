package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/services"
	"goapi/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GameList(w http.ResponseWriter, r *http.Request) {
	rawg := services.RawgService()
	resp, err := rawg.SearchGame(r.URL.Query())

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	db.OpenCollection("game-list")

	w.Write(jsonResponse)
}

func GameDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	rawg := services.RawgService()

	gameDetail, err := rawg.GetGameDetail(slug)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			utils.HandleApiErrors(w, http.StatusNotFound, "")
			return
		}
		fmt.Printf("ERROR: %v\n", err)
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	jsonResponse, err := json.Marshal(gameDetail)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	w.Write(jsonResponse)
}
