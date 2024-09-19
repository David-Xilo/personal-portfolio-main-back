package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
)

type NewsWithTopic struct {
	NewsList         []*models.NewsDTO           `json:"news_list"`
	TopicOfTheSeason *models.TopicOfTheSeasonDTO `json:"topic_of_the_season"`
}

func GetNewsByGenre(c *gin.Context, genre models.NewsGenres, db *gorm.DB) {
	newsList := getNews(genre, db)
	c.JSON(http.StatusOK, gin.H{"message": newsList})
}

func getNews(genre models.NewsGenres, db *gorm.DB) []*models.NewsDTO {
	var newsList []*models.News

	if err := db.
		Where("genre = ?", genre).
		Order("created_at desc").
		Limit(5).
		Find(&newsList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.NewsDTO{}
		}
		panic(err)
	}

	newsListDTO := models.ToNewsListDTO(newsList)
	return newsListDTO
}

func GetTopicOfTheSeasonByGenre(c *gin.Context, genre models.NewsGenres, db *gorm.DB) {
	result, err := getTopNewsWithTopic(genre, db)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func getTopNewsWithTopic(genre models.NewsGenres, db *gorm.DB) (*NewsWithTopic, error) {
	var newsList []*models.News
	var topicOfTheSeason *models.TopicOfTheSeasons

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

	newsListDTO := models.ToNewsListDTO(newsList)
	topicOfTheSeasonDTO := models.ToTopicOfTheSeasonDTO(topicOfTheSeason)

	return &NewsWithTopic{
		NewsList:         newsListDTO,
		TopicOfTheSeason: topicOfTheSeasonDTO,
	}, nil
}
