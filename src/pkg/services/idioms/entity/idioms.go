package entity

import "time"

type Idioms struct {
	TrimedIdiom string
	Idiom       string
	Meaning     string
	Level       IdiomLevel
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type IdiomLevel uint64

const (
	A1 IdiomLevel = iota + 1
	A2
	B1
	B2
	C1
)
