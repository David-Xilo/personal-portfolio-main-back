package models

import (
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
	ID             uint           `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      *time.Time     `json:"deleted_at,omitempty"`
	Topic          string         `json:"topic"`
	Genre          NewsGenres     `json:"genre"`
	TopicTimestamp time.Time      `json:"topic_timestamp"`
	Type           TimeframeTypes `json:"type"`
	CustomStart    *time.Time     `json:"custom_start"`
	CustomEnd      *time.Time     `json:"custom_end"`
}

type TopicOfTheSeasonDTO struct {
	Topic          string         `json:"topic"`
	Genre          NewsGenres     `json:"genre"`
	TopicTimestamp time.Time      `json:"topic_timestamp"`
	Type           TimeframeTypes `json:"type"`
	CustomStart    *time.Time     `json:"custom_start"`
	CustomEnd      *time.Time     `json:"custom_end"`
}

func ToTopicOfTheSeasonDTO(topicOfTheSeasons TopicOfTheSeasons) TopicOfTheSeasonDTO {
	return TopicOfTheSeasonDTO{
		Topic:          topicOfTheSeasons.Topic,
		Genre:          topicOfTheSeasons.Genre,
		TopicTimestamp: topicOfTheSeasons.TopicTimestamp,
		Type:           topicOfTheSeasons.Type,
		CustomStart:    topicOfTheSeasons.CustomStart,
		CustomEnd:      topicOfTheSeasons.CustomEnd,
	}
}
