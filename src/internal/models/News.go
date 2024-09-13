package models

import "gorm.io/gorm"

type NewsSentiments string

const (
	SentimentUndefined   NewsSentiments = "undefined"
	SentimentGood        NewsSentiments = "good"
	SentimentIndifferent NewsSentiments = "indifferent"
	SentimentBad         NewsSentiments = "bad"
)

type NewsGenres string

const (
	NewsGenreTech    NewsGenres = "tech"
	NewsGenreGaming  NewsGenres = "gaming"
	NewsGenreFinance NewsGenres = "finance"
)

type News struct {
	gorm.Model
	Headline     string         `json:"headline"`
	LinkToSource string         `json:"link-to-source"`
	Description  string         `json:"description"`
	Sentiment    NewsSentiments `json:"sentiment"`
	Genre        NewsGenres     `json:"genre"`
}
