package models

import "gorm.io/gorm"

type Contacts struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	LinkedIn string `json:"linked-in"`
	Github   string `json:"github"`
}
