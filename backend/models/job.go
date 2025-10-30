package models

import "time"

type Job struct {
	ID        int       `json:"id"`
	ImagePath string    `json:"image_path"`
	Operation string    `json:"operation"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
