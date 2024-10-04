package models

import "time"

type Video struct {
	Id           string    `json:"id" gorm:"primary_key"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
}
