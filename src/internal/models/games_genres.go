package models

type GameGenres string

const (
	GameGenreUndefined GameGenres = "undefined"
	GameGenreStrategy  GameGenres = "strategy"
	GameGenreTableTop  GameGenres = "table top"
	GameGenreRpg       GameGenres = "rpg"
)

func GetAllGameGenres() []GameGenres {
	return []GameGenres{
		GameGenreUndefined,
		GameGenreStrategy,
		GameGenreTableTop,
		GameGenreRpg,
	}
}
