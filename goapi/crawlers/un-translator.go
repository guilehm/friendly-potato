package crawlers

import (
	"context"
	"fmt"
	"goapi/db"
	"goapi/models"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var reportsCollection = db.OpenCollection("reports", "un")

func (c UNCrawler) Translate() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response models.CrawlResponse
	responseCollection.FindOne(
		ctx, bson.M{"translated": false},
	).Decode(&response)

	body := response.Body

	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		fmt.Println("could not parse body for", response.Url)
		return err
	}

	detailsXpath := `//div[@class="metadata-row"]/span[contains(text(), "%v")]/following-sibling::span`
	titleXpath := htmlquery.FindOne(doc, fmt.Sprintf(detailsXpath, "Title"))

	title := htmlquery.InnerText(titleXpath)
	report := models.UNReport{
		Url:   response.Url,
		Title: title,
	}

	symbolXpath := htmlquery.FindOne(doc, fmt.Sprintf(detailsXpath, "Symbol"))
	if symbolXpath != nil {
		symbol := htmlquery.InnerText(symbolXpath)
		report.Symbol = symbol
	}

	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	_, err = reportsCollection.UpdateOne(
		ctx, bson.M{"url": response.Url}, bson.D{{Key: "$set", Value: report}}, &opt,
	)

	if err != nil {
		fmt.Printf("An error ocurred while saving report for %s\n%s\n", response.Url, err)
		return err
	}

	return nil
}
