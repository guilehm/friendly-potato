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

type SearchResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}

type ratings struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
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
	Rating           float64     `json:"rating"`
	RatingTop        int         `json:"rating_top"`
	Ratings          []ratings   `json:"ratings"`
	RatingsCount     int         `json:"ratings_count"`
	ReviewsTextCount int         `json:"reviews_text_count"`
	Added            int         `json:"added"`
	Metacritic       int         `json:"metacritic"`
	Playtime         int         `json:"playtime"`
	SuggestionsCount int         `json:"suggestions_count"`
	Updated          string      `json:"updated"`
	EsrbRating       esrbRating  `json:"esrb_rating"`
	Platforms        []platforms `json:"platforms"`
}

func (r rawgService) SearchGame(query string) (SearchResponse, error) {
	endpoint, err := url.Parse(fmt.Sprintf("%v/%v", RAWG_API_URL, GAMES_ENDPOINT))
	if err != nil {
		return SearchResponse{}, err
	}

	q := endpoint.Query()
	q.Set("key", r.ApiKey)
	q.Set("search", query)
	endpoint.RawQuery = q.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil {
		return SearchResponse{}, err
	}
	var jsonResponse SearchResponse
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
