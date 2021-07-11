package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/models"
	"net/http"
	"net/url"
	"os"
)

const RAWG_API_URL = "https://api.rawg.io/api"
const GAMES_ENDPOINT = "games"

var ErrNotFound = errors.New("not found")

type rawgService struct {
	ApiKey string
}

func (r rawgService) GetGameDetail(gameSlug string) (models.GameStruct, error) {
	endpoint := fmt.Sprintf(
		"%v/%v/%v?key=%v", RAWG_API_URL, GAMES_ENDPOINT, gameSlug, r.ApiKey,
	)
	resp, err := http.Get(endpoint)
	if err != nil {
		return models.GameStruct{}, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return models.GameStruct{}, ErrNotFound
	}

	var jsonResponse models.GameStruct
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	return jsonResponse, err
}

func (r rawgService) SearchGame(queries url.Values) (models.SearchResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%v/%v", RAWG_API_URL, GAMES_ENDPOINT))
	if err != nil {
		return models.SearchResponse{}, err
	}

	q := endpoint.Query()
	q.Set("key", r.ApiKey)
	for key, values := range queries {
		for _, v := range values {
			q.Set(key, v)
		}
	}
	endpoint.RawQuery = q.Encode()
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return models.SearchResponse{}, err
	}
	var response models.SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func RawgService() rawgService {
	rawgApiKey := os.Getenv("RAWG_API_KEY")
	if rawgApiKey == "" {
		panic("RAWG_API_KEY not set")
	}
	service := rawgService{rawgApiKey}
	return service
}
