package models

import (
	"gorm.io/gorm"
)

type GameGenres string

const (
	GameGenreUndefined GameGenres = "undefined"
	GameGenreStrategy  GameGenres = "strategy"
	GameGenreTableTop  GameGenres = "table top"
)

type Games struct {
	gorm.Model
	Title       string     `json:"title"`
	Genre       GameGenres `json:"genre"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link_to_git"`
	LinkToStore string     `json:"link_to_store"`
}

func GetAllGameGenres() []GameGenres {
	return []GameGenres{
		GameGenreUndefined,
		GameGenreStrategy,
		GameGenreTableTop,
	}
}
