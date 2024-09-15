package models

import (
	"time"
)

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
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at,omitempty"`
	Headline     string         `json:"headline"`
	LinkToSource string         `json:"link-to-source"`
	Description  string         `json:"description"`
	Sentiment    NewsSentiments `json:"sentiment"`
	Genre        NewsGenres     `json:"genre"`
}
