package main

import "time"

type user struct {
	name  string
	feeds []feed
}

type feed struct {
	name  string
	url   string
	items []feeditem
}

type feeditem struct {
	title   string
	content string
	date    time.Time
	url     string
}
