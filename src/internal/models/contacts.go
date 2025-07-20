package models

import (
	"gorm.io/gorm"
	"time"
)

type Contacts struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;not null" json:"email"`
	Active    bool           `gorm:"not null" json:"active"`
	Linkedin  string         `gorm:"size:255" json:"linkedin"`
	Github    string         `gorm:"size:255" json:"github"`
	Credly    string         `gorm:"size:255" json:"credly"`
}

type ContactsDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	LinkedIn string `json:"linkedin"`
	Github   string `json:"github"`
	Credly   string `json:"credly"`
}

func ToContactsDTO(contact *Contacts) *ContactsDTO {
	return &ContactsDTO{
		Name:     contact.Name,
		Email:    contact.Email,
		LinkedIn: contact.Linkedin,
		Github:   contact.Github,
		Credly:   contact.Credly,
	}
}
