package models

type ProjectTypes string

const (
	ProjectTypeUndefined ProjectTypes = "undefined"
	ProjectTypeTech      ProjectTypes = "tech"
	ProjectTypeGame      ProjectTypes = "game"
	ProjectTypeFinance   ProjectTypes = "finance"
)

func GetAllProjectTypes() []ProjectTypes {
	return []ProjectTypes{
		ProjectTypeUndefined,
		ProjectTypeTech,
		ProjectTypeGame,
		ProjectTypeFinance,
	}
}
