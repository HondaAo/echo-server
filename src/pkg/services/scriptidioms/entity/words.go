package entity

import (
	"time"
)

type ScriptIdioms struct {
	Idiom     string
	VideoID   string
	ScriptID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
