package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type date time.Time

func (d *date) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.RFC1123Z, string(data))

	*d = date(t)
	return err
}

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title string `xml:"title"`
	Desc  string `xml:"description"`

	Items []NyaaTorrentItem `xml:"item"`
}

type NyaaTorrentItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate date   `xml:"pubDate"`
	Seeder  int    `xml:"seeders"`
	Leecher int    `xml:"leechers"`
}

func ParseTorrentItem(input string) (item NyaaTorrentItem, err error) {
	err = xml.Unmarshal([]byte(input), &item)

	return
}

func ParseRSS(input string) (rss RSS, err error) {
	err = xml.Unmarshal([]byte(input), &rss)

	return
}

func FetchTorrentItem(url string) (items []NyaaTorrentItem, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	rss, err := ParseRSS(string(body))

	if err == nil {
		items = rss.Channel.Items
	}

	return
}
