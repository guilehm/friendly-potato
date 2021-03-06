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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// crawler.Crawl(10)

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
	err = xml.Unmarshal(unzipData, &sitemap)
	if err != nil {
		return err
	}

	for _, sitemap := range sitemap.Sitemaps {
		resp, err := c.GetResponse(sitemap.Location)
		if err != nil {
			return err
		}

		gReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}

		var sitemapDetail models.SitemapDetail

		unzipData, err := ioutil.ReadAll(gReader)
		if err != nil {
			return err
		}
		_ = resp.Body.Close()

		err = xml.Unmarshal(unzipData, &sitemapDetail)
		if err != nil {
			return err
		}
		sitemapCount := len(sitemapDetail.Sitemaps)
		docs := make([]interface{}, sitemapCount)
		for i, s := range sitemapDetail.Sitemaps {
			s.ID = primitive.NewObjectID()
			docs[i] = s
		}

		opts := options.InsertMany().SetOrdered(false)
		_, err = sitemapsCollection.InsertMany(
			context.Background(), docs, opts,
		)
		if err != nil {
			return err
		}
		fmt.Printf("Success saving %v sitemaps\n", sitemapCount)
	}

	return nil
}

func (c UNCrawler) SaveBodyData(url string) error {

	resp, err := c.GetResponse(url)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	body := buf.String()
	response := models.CrawlResponse{
		Url:        url,
		StatusCode: resp.StatusCode,
		Body:       body,
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

func (c UNCrawler) Crawl(limit, skip int64) error {
	ctx := context.Background()
	opts := options.Find().SetLimit(limit).SetSkip(skip)
	cur, err := sitemapsCollection.Find(
		ctx,
		bson.M{"crawled": false},
		opts,
	)

	if err != nil {
		return err
	}

	for cur.Next(ctx) {

		var sitemap models.Sitemap
		err := cur.Decode(&sitemap)
		if err != nil {
			return err
		}

		if err := c.SaveBodyData(sitemap.Location); err != nil {
			fmt.Println("Error while saving body data for", sitemap.Location)
		}

		fmt.Println("\t	Updating ", sitemap.Location)
		_, err = sitemapsCollection.UpdateOne(
			ctx,
			bson.M{"_id": sitemap.ID},
			bson.D{
				{Key: "$set", Value: bson.D{{Key: "crawled", Value: true}}},
			},
		)

		if err != nil {
			fmt.Printf(
				"An error ocurred while updating %s\n\t%s\n",
				sitemap.Location,
				err,
			)
		}
	}
	cur.Close(ctx)

	return nil
}
