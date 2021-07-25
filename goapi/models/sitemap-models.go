package models

import "encoding/xml"

type Sitemap struct {
	Location string `xml:"loc"`
	LastMod  string `xml:"lastmod"`
}

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}
