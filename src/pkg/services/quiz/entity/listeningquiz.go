package entity

import "time"

type ListeningQuiz struct {
	VideoID   string
	ScriptID  uint64
	Word      string
	Level     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
