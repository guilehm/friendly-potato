package models

import "encoding/xml"

type Sitemap struct {
	Location string `xml:"loc" bson:"location"`
	LastMod  string `xml:"lastmod" bson:"lastmod"`
}

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex" bson:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap" bson:"sitemap"`
}
