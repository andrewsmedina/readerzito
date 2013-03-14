package main

import (
	"encoding/xml"
	"labix.org/v2/mgo"
	"time"
)

var collection *mgo.Collection

type user struct {
	name  string
	feeds []feed
}

type feed struct {
	Name  string
	Url   string
	items []feeditem
}

type feeditem struct {
	title   string
	content string
	date    time.Time
	url     string
}

// subscribe stores a feed.
func subscribe(name, url string) error {
	f := feed{Name: name, Url: url}
	return collection.Insert(f)
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

func parse(data string) ([]feeditem, error) {
	var r RSS
	var items []feeditem
	err := xml.Unmarshal([]byte(data), &r)
	if err != nil {
		return items, err
	}
	for _, item := range r.Channel.Item {
		i := feeditem{
			title:   item.Title,
			content: item.Description,
			url:     item.Link,
		}
		items = append(items, i)
	}
	return items, nil
}
