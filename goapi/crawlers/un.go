package crawlers

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

	resourcesPath := "./resources/" + c.Name
	err = os.MkdirAll(resourcesPath, os.ModePerm)
	if err != nil {
		return err
	}

	writer, err := os.Create(fmt.Sprintf("%v/%v", resourcesPath, filename))
	if err != nil {
		return err
	}
	defer writer.Close()

	if _, err = io.Copy(writer, gReader); err != nil {
		return err
	}

	return nil
}
