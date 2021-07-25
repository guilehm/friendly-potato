package crawlers

import (
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"goapi/models"
	"io/ioutil"
	"net/http"
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

	var xmlData models.SitemapIndex

	sitemap, err := ioutil.ReadAll(gReader)
	if err != nil {
		return err
	}
	xml.Unmarshal(sitemap, &xmlData)

	for _, sitemap := range xmlData.Sitemaps {
		fmt.Println(sitemap.Location)
	}

	return nil
}
