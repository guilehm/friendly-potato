package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const RAWG_API_URL = "https://api.rawg.io/api"
const GAMES_ENDPOINT = "games"

type rawgService struct {
	ApiKey string
}

type searchResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}
type esrbRating struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
type platform struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
type requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}
type platforms struct {
	Platform     platform     `json:"platform"`
	ReleasedAt   string       `json:"released_at"`
	Requirements requirements `json:"requirements"`
}
type results struct {
	ID               int         `json:"id"`
	Slug             string      `json:"slug"`
	Name             string      `json:"name"`
	Released         string      `json:"released"`
	Tba              bool        `json:"tba"`
	BackgroundImage  string      `json:"background_image"`
	Rating           int         `json:"rating"`
	RatingTop        int         `json:"rating_top"`
	RatingsCount     int         `json:"ratings_count"`
	ReviewsTextCount string      `json:"reviews_text_count"`
	Added            int         `json:"added"`
	Metacritic       int         `json:"metacritic"`
	Playtime         int         `json:"playtime"`
	SuggestionsCount int         `json:"suggestions_count"`
	Updated          string      `json:"updated"`
	EsrbRating       esrbRating  `json:"esrb_rating"`
	Platforms        []platforms `json:"platforms"`
}

func (r rawgService) SearchGame(query string) (searchResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%v/%v", RAWG_API_URL, GAMES_ENDPOINT))
	if err != nil {
		return searchResponse{}, err
	}

	q := endpoint.Query()
	q.Set("key", r.ApiKey)
	endpoint.RawQuery = q.Encode()

	fmt.Printf("%v\n", endpoint)
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return searchResponse{}, err
	}
	var jsonResponse searchResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	return jsonResponse, err
}

func RawgService() rawgService {
	rawgApiKey := os.Getenv("RAWG_API_KEY")
	if rawgApiKey == "" {
		panic("RAWG_API_KEY not set")
	}
	service := rawgService{rawgApiKey}
	return service
}
