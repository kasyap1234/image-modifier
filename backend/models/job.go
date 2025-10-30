package models

import "time"

type ImageConfig struct {
	Config MinioConfig `json:"config"`
}
type Job struct {
	ID        int       `json:"id"`
	Image ImageConfig `json:"image_config"`
	Operation string    `json:"operation"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}



