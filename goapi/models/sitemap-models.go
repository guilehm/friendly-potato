package models

import "encoding/xml"

type Sitemap struct {
	Location   string `xml:"loc" bson:"location"`
	LastMod    string `xml:"lastmod" bson:"lastmod"`
	ChangeFreq string `xml:"changefreq" bson:"changefreq,omitempty"`
	Priority   string `xml:"priority" bson:"priority,omitempty"`
	Crawled    bool   `xml:"-" bson:"crawled"`
}

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex" bson:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap" bson:"sitemap"`
}

type SitemapDetail struct {
	XMLName  xml.Name  `xml:"urlset" bson:"urlset"`
	Sitemaps []Sitemap `xml:"url" bson:"url"`
}
