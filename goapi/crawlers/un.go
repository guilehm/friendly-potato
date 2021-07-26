package crawlers

import (
	"compress/gzip"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/models"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type UNCrawler struct {
	Name        string
	BaseUrl     string
	SiteMapPath string
}

func (c UNCrawler) GetAllUrlsFromSitemap() error {

	url := c.BaseUrl + c.SiteMapPath

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP error: " + resp.Status)
	}

	gReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	var sitemap models.SitemapIndex

	unzipData, err := ioutil.ReadAll(gReader)
	if err != nil {
		return err
	}
	xml.Unmarshal(unzipData, &sitemap)

	var unUrlsCollection = db.OpenCollection("urls", "un")

	for _, sitemap := range sitemap.Sitemaps {
		fmt.Println("Requesting", sitemap.Location)
		resp, err := http.Get(sitemap.Location)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New("HTTP error: " + resp.Status)
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
		xml.Unmarshal(unzipData, &sitemapDetail)
		opts := options.InsertMany().SetOrdered(false)
		docs := make([]interface{}, len(sitemapDetail.Sitemaps))
		for i, v := range sitemapDetail.Sitemaps {
			docs[i] = v
		}
		_, err = unUrlsCollection.InsertMany(
			context.TODO(), docs, opts,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Success saving %v\n", len(sitemapDetail.Sitemaps))
	}

	return nil
}
