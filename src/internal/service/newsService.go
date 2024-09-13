package service

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
)

type NewsWithTopic struct {
	NewsList         []*models.News            `json:"news_list"`
	TopicOfTheSeason *models.TopicOfTheSeasons `json:"topic_of_the_season"`
}

func GetNewsByGenre(w http.ResponseWriter, genre models.NewsGenres, db *gorm.DB) {
	newsList := getNews(genre, db)
	GetJSONData(w, newsList)
}

func getNews(genre models.NewsGenres, db *gorm.DB) []*models.News {
	var newsList []*models.News

	if err := db.
		Where("genre = ?", genre).
		Order("created_at desc").
		Limit(5).
		Find(&newsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.News{}
		}
		panic(err)
	}

	return newsList // Return the list of news if found
}

func GetTopicOfTheSeasonByGenre(w http.ResponseWriter, genre models.NewsGenres, db *gorm.DB) {
	result, err := getTopNewsWithTopic(genre, db)
	if err != nil {
		panic(err)
	}
	GetJSONData(w, result)
}

func getTopNewsWithTopic(genre models.NewsGenres, db *gorm.DB) (*NewsWithTopic, error) {
	var newsList []*models.News
	var topicOfTheSeason models.TopicOfTheSeasons

	if err := db.Where("genre = ?", genre).
		Order("created_at desc").
		First(&topicOfTheSeason).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		panic(err)
	}

	if err := db.Table("news_topic_of_the_seasons").
		Select("news.*").
		Joins("JOIN news ON news_topic_of_the_seasons.news_id = news.id").
		Where("news_topic_of_the_seasons.topic_of_the_season_id = ?", topicOfTheSeason.ID).
		Order("news_topic_of_the_seasons.created_at desc").
		Limit(5).
		Scan(&newsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		panic(err)
	}

	return &NewsWithTopic{
		NewsList:         newsList,
		TopicOfTheSeason: &topicOfTheSeason,
	}, nil
}
