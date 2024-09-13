package models

import (
	"gorm.io/gorm"
	"time"
)

type TimeframeTypes string

const (
	TimeframeCustom TimeframeTypes = "custom"
	Timeframe1D     TimeframeTypes = "1d"
	Timeframe1W     TimeframeTypes = "1w"
	Timeframe1M     TimeframeTypes = "1m"
	Timeframe3M     TimeframeTypes = "3m"
	Timeframe6M     TimeframeTypes = "6m"
	Timeframe1Y     TimeframeTypes = "1y"
)

type TopicOfTheSeasons struct {
	gorm.Model
	Topic          string         `json:"topic"`
	Genre          NewsGenres     `json:"genre"`
	TopicTimestamp time.Time      `json:"topic_timestamp"`
	Type           TimeframeTypes `json:"type"`
	CustomStart    *time.Time     `json:"custom_start"`
	CustomEnd      *time.Time     `json:"custom_end"`
}
