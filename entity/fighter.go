package entity

import "time"

type Fighter struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Results     string    `json:"results"`
	Description string    `json:"description"`
	Backbone    string    `json:"backbone"`
	Age         int       `json:"age"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
