package models

type CrawlResponse struct {
	Url        string `bson:"url"`
	StatusCode int    `bson:"status_code"`
	Error      string `bson:"error"`
	Body       string `bson:"body"`
	Translated bool   `bson:"translated"`
}

type UNReport struct {
	Url         string `bson:"url"`
	Title       string `bson:"title"`
	Symbol      string `bson:"symbol"`
	Imprint     string `bson:"imprint"`
	Description string `bson:"description"`
}
