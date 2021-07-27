package crawlers

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/models"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// START CRAWLER
// crawler := UNCrawler{
// 	Name:        "un",
// 	BaseUrl:     "https://digitallibrary.un.org",
// 	SiteMapPath: "/sitemap_index.xml.gz",
// }

// err := crawler.GetAllUrlsFromSitemaps()
// if err != nil {
// 	fmt.Println("Error while trying to get sitemap", err)
// }

var responseCollection = db.OpenCollection("response", "un")
var sitemapsCollection = db.OpenCollection("sitemaps", "un")

type UNCrawler struct {
	Name        string
	BaseUrl     string
	SiteMapPath string
}

func (c UNCrawler) GetResponse(url string) (*http.Response, error) {
	fmt.Println("Requesting", url)
	resp, err := http.Get(url)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("HTTP error: " + resp.Status)
	}

	return resp, nil
}

func (c UNCrawler) GetAllUrlsFromSitemaps() error {

	url := c.BaseUrl + c.SiteMapPath

	resp, err := c.GetResponse(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	unzipData, err := ioutil.ReadAll(gReader)
	if err != nil {
		return err
	}

	var sitemap models.SitemapIndex
	xml.Unmarshal(unzipData, &sitemap)

	for _, sitemap := range sitemap.Sitemaps {
		fmt.Println("Requesting", sitemap.Location)
		resp, err := c.GetResponse(sitemap.Location)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		gReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}

		var sitemapDetail models.SitemapDetail

		unzipData, err := ioutil.ReadAll(gReader)
		if err != nil {
			return err
		}
		xml.Unmarshal(unzipData, &sitemapDetail)
		sitemapCount := len(sitemapDetail.Sitemaps)
		docs := make([]interface{}, sitemapCount)
		for i, v := range sitemapDetail.Sitemaps {
			docs[i] = v
		}

		opts := options.InsertMany().SetOrdered(false)
		_, err = sitemapsCollection.InsertMany(
			context.TODO(), docs, opts,
		)
		if err != nil {
			return err
		}
		fmt.Printf("Success saving %v\n", sitemapCount)
	}

	return nil
}

func (c UNCrawler) SaveBodyData(url string) error {

	resp, err := c.GetResponse(url)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()
	response := models.CrawlResponse{
		Url:        url,
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}
	if err != nil {
		response.Error = err.Error()
	}
	defer resp.Body.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = responseCollection.InsertOne(ctx, response)
	return err

}

func (c UNCrawler) Crawl(limit int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Find().SetLimit(limit)
	cur, err := sitemapsCollection.Find(
		ctx,
		bson.M{"crawled": false},
		opts,
	)
	defer cancel()

	if err != nil {
		return err
	}

	for cur.Next(context.Background()) {

		var sitemap models.Sitemap
		err := cur.Decode(&sitemap)
		if err != nil {
			return err
		}

		if err := c.SaveBodyData(sitemap.Location); err != nil {
			fmt.Println("Error while saving body data for", sitemap.Location)
		}

		fmt.Println("Trying to update", sitemap.Location)
		_, err = sitemapsCollection.UpdateOne(
			ctx,
			bson.M{"location": sitemap.Location},
			bson.D{
				{"$set", bson.D{{"crawled", true}}},
			},
		)

		if err != nil {
			fmt.Printf("An error ocurred while updating %s", sitemap.Location)
		}
	}

	cur.Close(ctx)

	return nil
}
