package entity

import "time"

type WordQuiz struct {
	QuizID    uint64
	Word      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
