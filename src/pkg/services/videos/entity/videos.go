package entity

import (
	"time"
)

type Video struct {
	VideoID    string
	Title      string
	URL        string
	Start      float64
	End        float64
	CategoryID CategoryID
	Level      VideoLevel
	Display    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type VideoLevel uint64

type CategoryID uint64

const (
	A1 VideoLevel = iota + 1
	A2
	B1
	B2
	C1
)

const (
	Drama CategoryID = iota + 1
	Anime
	Movie
	Politics
	Economy
	Geograpy
)

type SearchCondition struct {
	VideoIDs   []string
	Limit      uint64
	Level      uint64
	CategoryID uint64
}

const DEFAULT_MAX_LIMIT = 100
const DEFAULT_CATEGORY_ID = 0
const DEFAULT_LEVEL = 0
