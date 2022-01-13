package main

import "gorm.io/gorm"

type Response struct {
	gorm.Model
	OriginalUrl  string `json:"originalUrl"`
	ShortenedUrl string `json:"shortenedUrl"`
}

type OriginalUrl struct {
	url string `json:"url"`
}

var urls []Response
