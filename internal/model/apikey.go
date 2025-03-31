package model

import "time"

type ApiKey struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"unique"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}
