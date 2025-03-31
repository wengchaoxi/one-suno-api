package model

import "time"

type Provider struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}
