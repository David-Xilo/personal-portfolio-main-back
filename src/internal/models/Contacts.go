package models

import (
	"time"
)

type Contacts struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	LinkedIn  string     `json:"linked-in"`
	Github    string     `json:"github"`
}
