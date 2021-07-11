package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/services"
	"goapi/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GameList(w http.ResponseWriter, r *http.Request) {
	rawg := services.RawgService()
	resp, err := rawg.SearchGame(r.URL.Query())

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
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
		fmt.Println("Error: invalid ID")
		utils.HandleApiErrors(w, http.StatusBadRequest, "invalid ID")
	}

	resp, err := rawg.GetGameDetail(intId)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			utils.HandleApiErrors(w, http.StatusNotFound, "")
			return
		}
		fmt.Printf("ERROR: %v\n", err)
		utils.HandleApiErrors(w, http.StatusInternalServerError, "")
		return
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	w.Write(jsonResponse)
}
