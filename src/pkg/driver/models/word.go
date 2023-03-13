package models

import "time"

type Words struct {
	WordID    uint64 `gorm:"primaryKey"`
	Word      string
	Mean      string
	Comment   string
	Level     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
