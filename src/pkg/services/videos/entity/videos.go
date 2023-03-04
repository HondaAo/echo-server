package entity

import (
	"time"
)

type Video struct {
	VideoID   string
	Title     string
	URL       string
	Start     float64
	End       float64
	Level     VideoLevel
	Display   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VideoLevel uint64

const (
	A1 VideoLevel = iota + 1
	A2
	B1
	B2
	C1
)
