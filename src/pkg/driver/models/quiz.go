package models

import (
	"time"
)

type ListeningQuizFillBlank struct {
	VideoID   string `gorm:"primarykey"`
	ScriptID  uint64 `gorm:"primaryKey"`
	Word      string `gorm:"primaryKey"`
	Level     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListeningQuizDiscription struct {
	QuizID    uint64 `gorm:"primaryKey"`
	VideoID   string
	Level     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Choices struct {
	QuizID    uint64 `gorm:"primaryKey"`
	Choice1   string
	Choice2   string
	Choice3   string
	Choice4   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WritingQuiz struct {
	VideoID   string
	ScriptID  uint64
	Level     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WordQuiz struct {
	QuizID    uint64 `gorm:"primaryKey"`
	Word      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
