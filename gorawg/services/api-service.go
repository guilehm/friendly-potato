package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const RAWG_API_URL = "https://api.rawg.io/api"
const GAMES_ENDPOINT = "games"

type RawgService struct {
	ApiKey string
}

type searchResponse struct {
	Count int `json:"count"`
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func (r RawgService) SearchGame(query string, target interface{}) *searchResponse {
	endpoint := fmt.Sprintf("%v/%v", RAWG_API_URL, GAMES_ENDPOINT)
	sr := new(searchResponse)
	getJson(endpoint, sr)
	return sr
}
