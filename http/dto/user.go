package dto

import "time"

type User struct {
	Id        int       `json:"user_id,omitempty"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
