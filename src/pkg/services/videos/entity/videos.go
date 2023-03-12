package entity

import (
	"time"
)

type Video struct {
	VideoID     string
	Title       string
	URL         string
	Start       float64
	End         float64
	CategoryID  CategoryID
	Level       VideoLevel
	Display     bool
	EnglishType EnglishType
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	News
	Economy
	Geograpy
	TVshow
	Travel
	Sports
	Youtuber
)

type SearchCondition struct {
	VideoIDs    []string
	Limit       uint64
	Levels      []uint64
	CategoryIDs []uint64
	Date        time.Time
	Length      []uint64
	TypeIDs     []uint8
}

const DEFAULT_MAX_LIMIT = 100
const DEFAULT_CATEGORY_ID = 0
const DEFAULT_LEVEL = 0

type EnglishType uint8

const (
	America EnglishType = iota
	England
	Other
)

func GetDate(dateId uint8) time.Time {
	t := time.Now()
	switch dateId {
	case 1:
		return t.AddDate(0, 0, -7)
	case 2:
		return t.AddDate(0, -1, 0)
	case 3:
		return t.AddDate(-1, 0, 0)
	default:
		return t.AddDate(-2, 0, 0)
	}
}

func GetLength(lengthId uint8) []uint64 {
	switch lengthId {
	case 1:
		return []uint64{0, 59}
	case 2:
		return []uint64{60, 179}
	case 3:
		return []uint64{180, 6000}
	default:
		return []uint64{0, 6000}
	}
}
