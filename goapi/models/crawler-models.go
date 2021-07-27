package models

type CrawlResponse struct {
	Url        string `bson:"url"`
	StatusCode int    `bson:"status_code"`
	Error      string `bson:"error"`
	Body       string `bson:"body"`
}
