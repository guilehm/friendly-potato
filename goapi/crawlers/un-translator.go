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

func (c UNCrawler) TranslateMany(limit int64) error {
	ctx := context.Background()
	opts := options.Find().SetLimit(limit).SetProjection(
		bson.M{"_id": 0, "url": 1},
	)

	cur, err := responseCollection.Find(
		ctx,
		bson.M{"translated": false},
		opts,
	)

	if err != nil {
		fmt.Println("could not retrieve data to make translations")
		return err
	}

	for cur.Next(ctx) {
		var response models.CrawlResponse
		err := cur.Decode(&response)

		if err != nil {
			fmt.Println("could not decode sitemap response")
			return err
		}

		c.Translate(response.Url)
	}

	return nil
}

func (c UNCrawler) Translate(url string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response models.CrawlResponse
	responseCollection.FindOne(
		ctx, bson.M{"url": url},
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

	upsert = false
	_, err = responseCollection.UpdateOne(
		ctx, bson.M{"url": response.Url}, bson.D{{Key: "$set", Value: bson.M{"translated": true}}}, &opt,
	)

	if err != nil {
		fmt.Printf("An error ocurred while updating response for %s\n%s\n", response.Url, err)
		return err
	}

	return nil
}
