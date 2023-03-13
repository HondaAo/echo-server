package entity

import "time"

type Answer struct {
	QuizID    uint64
	Choice1   string
	Choice2   string
	Choice3   string
	Choice4   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
