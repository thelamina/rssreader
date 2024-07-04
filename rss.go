package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	http_client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := http_client.Get(url)

	if err != nil {
		return RSSFeed{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFeed{}, err
	}
	rss_feed := RSSFeed{}
	err = xml.Unmarshal(data, &rss_feed)

	if err != nil {
		return RSSFeed{}, err
	}

	return rss_feed, nil

}
