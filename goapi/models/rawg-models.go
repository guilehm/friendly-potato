package models

type GameStruct struct {
	ID                        int                   `json:"id"`
	Slug                      string                `json:"slug"`
	Name                      string                `json:"name"`
	NameOriginal              string                `json:"name_original"`
	Description               string                `json:"description"`
	Metacritic                int                   `json:"metacritic"`
	MetacriticPlatforms       []MetacriticPlatforms `json:"metacritic_platforms"`
	Released                  string                `json:"released"`
	Tba                       bool                  `json:"tba"`
	Updated                   string                `json:"updated"`
	BackgroundImage           string                `json:"background_image"`
	BackgroundImageAdditional string                `json:"background_image_additional"`
	Website                   string                `json:"website"`
	Rating                    float64               `json:"rating"`
	RatingTop                 int                   `json:"rating_top"`
	Ratings                   []Ratings             `json:"ratings"`
	Added                     int                   `json:"added"`
	AddedByStatus             AddedByStatus         `json:"added_by_status"`
	Playtime                  int                   `json:"playtime"`
	ScreenshotsCount          int                   `json:"screenshots_count"`
	MoviesCount               int                   `json:"movies_count"`
	CreatorsCount             int                   `json:"creators_count"`
	AchievementsCount         int                   `json:"achievements_count"`
	ParentAchievementsCount   int                   `json:"parent_achievements_count"`
	RedditURL                 string                `json:"reddit_url"`
	RedditName                string                `json:"reddit_name"`
	RedditDescription         string                `json:"reddit_description"`
	RedditLogo                string                `json:"reddit_logo"`
	RedditCount               int                   `json:"reddit_count"`
	TwitchCount               int                   `json:"twitch_count"`
	YoutubeCount              int                   `json:"youtube_count"`
	ReviewsTextCount          int                   `json:"reviews_text_count"`
	RatingsCount              int                   `json:"ratings_count"`
	SuggestionsCount          int                   `json:"suggestions_count"`
	AlternativeNames          []string              `json:"alternative_names"`
	MetacriticURL             string                `json:"metacritic_url"`
	ParentsCount              int                   `json:"parents_count"`
	AdditionsCount            int                   `json:"additions_count"`
	GameSeriesCount           int                   `json:"game_series_count"`
	UserGame                  interface{}           `json:"user_game"`
	ReviewsCount              int                   `json:"reviews_count"`
	SaturatedColor            string                `json:"saturated_color"`
	DominantColor             string                `json:"dominant_color"`
	ParentPlatforms           []ParentPlatforms     `json:"parent_platforms"`
	Platforms                 []Platforms           `json:"platforms"`
	Stores                    []Stores              `json:"stores"`
	Developers                []Developers          `json:"developers"`
	Genres                    []Genres              `json:"genres"`
	Tags                      []Tags                `json:"tags"`
	Publishers                []Publishers          `json:"publishers"`
	EsrbRating                EsrbRating            `json:"esrb_rating"`
	Clip                      interface{}           `json:"clip"`
	DescriptionRaw            string                `json:"description_raw"`
}
type MetacriticPlatforms struct {
	Metascore int      `json:"metascore"`
	URL       string   `json:"url"`
	Platform  Platform `json:"platform"`
}
type Ratings struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}
type AddedByStatus struct {
	Yet     int `json:"yet"`
	Owned   int `json:"owned"`
	Beaten  int `json:"beaten"`
	Toplay  int `json:"toplay"`
	Dropped int `json:"dropped"`
	Playing int `json:"playing"`
}
type ParentPlatforms struct {
	Platform Platform `json:"platform"`
}
type Platform struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Slug            string      `json:"slug"`
	Image           interface{} `json:"image,omitempty"`
	YearEnd         interface{} `json:"year_end,omitempty"`
	YearStart       int         `json:"year_start,omitempty"`
	GamesCount      int         `json:"games_count,omitempty"`
	ImageBackground string      `json:"image_background,omitempty"`
}
type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}
type Platforms struct {
	Platform     Platform     `json:"platform"`
	ReleasedAt   string       `json:"released_at"`
	Requirements Requirements `json:"requirements,omitempty"`
}
type Store struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Domain          string `json:"domain"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}
type Stores struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Store Store  `json:"store"`
}
type Developers struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}
type Genres struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}
type Tags struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Language        string `json:"language"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}
type Publishers struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}
type EsrbRating struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type SearchResponse struct {
	Count int `json:"count"`
	// Next     string    `json:"next"`
	// Previous string    `json:"previous"`
	Results []results `json:"results"`
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
