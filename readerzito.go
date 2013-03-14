package main

import (
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
