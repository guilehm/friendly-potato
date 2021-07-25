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

func (c UNCrawler) GetSiteMap(filename string) error {

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

	for _, sitemap := range sitemap.Sitemaps {
		fmt.Println(sitemap.Location)
	}

	var sitemapsCollection = db.OpenCollection("sitemaps")

	opts := options.InsertMany().SetOrdered(false)
	docs := make([]interface{}, len(sitemap.Sitemaps))
	for i, v := range sitemap.Sitemaps {
		docs[i] = v
	}

	results, err := sitemapsCollection.InsertMany(
		context.TODO(), docs, opts,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", results)

	return nil
}
