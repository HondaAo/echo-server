package entity

import "time"

type Word struct {
	WordID    uint64
	Word      string
	Mean      string
	Comment   string
	Level     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
