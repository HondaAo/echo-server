package models

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
)

type Idioms struct {
	TrimedIdiom string    `gorm:"trimed_idiom;primaryKey"`
	Idiom       string    `gorm:"idiom"`
	Meaning     string    `gorm:"meaning"`
	Level       uint64    `gorm:"level"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

func (i *Idioms) TableName() string {
	return "idioms"
}

func NewIdiomEntity(idioms *Idioms) *entity.Idioms {
	return &entity.Idioms{
		TrimedIdiom: idioms.TrimedIdiom,
		Idiom:       idioms.Idiom,
		Meaning:     idioms.Meaning,
		Level:       entity.IdiomLevel(idioms.Level),
		CreatedAt:   idioms.CreatedAt,
		UpdatedAt:   idioms.UpdatedAt,
	}
}
