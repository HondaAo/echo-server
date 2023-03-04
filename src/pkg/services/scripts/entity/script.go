package entity

import "time"

type Script struct {
	VideoID   string
	ScriptID  uint64
	Text      string
	Japanese  string
	TimeStamp float64
	Highlight bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
