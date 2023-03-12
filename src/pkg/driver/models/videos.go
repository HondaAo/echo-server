package models

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
)

type Video struct {
	VideoID    string `gorm:"primaryKey"`
	Title      string
	URL        string
	Start      float64
	End        float64
	Type       uint64    `gorm:"default:1"`
	Length     float64   `gorm:"default:60"`
	CategoryID uint64    `gorm:"default:1"`
	Level      uint64    `gorm:"level"`
	Display    bool      `gorm:"display"`
	CreatedAt  time.Time `gorm:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at"`
}

func (v Video) TableName() string {
	return "videos"
}

func NewEntity(video *Video) *entity.Video {
	return &entity.Video{
		VideoID:    video.VideoID,
		Title:      video.Title,
		URL:        video.URL,
		Start:      video.Start,
		End:        video.End,
		CategoryID: entity.CategoryID(video.CategoryID),
		Level:      entity.VideoLevel(video.Level),
		Display:    video.Display,
		CreatedAt:  video.CreatedAt,
		UpdatedAt:  video.UpdatedAt,
	}
}
