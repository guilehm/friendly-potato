package models

import "encoding/xml"

type Sitemap struct {
	Location   string `xml:"loc" bson:"location"`
	LastMod    string `xml:"lastmod" bson:"lastmod"`
	ChangeFreq string `xml:"changefreq" bson:"changefreq,omitempty"`
	Priority   string `xml:"priority" bson:"priority,omitempty"`
}

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex" bson:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap" bson:"sitemap"`
}
