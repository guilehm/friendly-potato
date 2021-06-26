package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const RAWG_API_URL = "https://api.rawg.io/api"
const GAMES_ENDPOINT = "games"

type rawgService struct {
	ApiKey string
}

type SearchResponse struct {
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

func (r rawgService) SearchGame(query string, target interface{}) SearchResponse {
	endpoint := fmt.Sprintf("%v/%v", RAWG_API_URL, GAMES_ENDPOINT)
	sr := new(SearchResponse)
	getJson(endpoint, sr)
	return SearchResponse{sr.Count}
}

func RawgService() rawgService {
	rawgApiKey := os.Getenv("RAWG_API_KEY")
	if rawgApiKey == "" {
		panic("RAWG_API_KEY not set")
	}
	service := rawgService{rawgApiKey}
	return service
}
