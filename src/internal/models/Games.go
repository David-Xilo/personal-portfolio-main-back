package models

import (
	"time"
)

type GameGenres string

const (
	GameGenreUndefined GameGenres = "undefined"
	GameGenreStrategy  GameGenres = "strategy"
	GameGenreTableTop  GameGenres = "table top"
)

type Games struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
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
