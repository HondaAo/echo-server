package models

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
)

type ScriptIdioms struct {
	VideoID   string    `gorm:"video_id;primaryKey"`
	ScriptID  uint64    `gorm:"script_id;primaryKey"`
	Word      string    `gorm:"word;primaryKey"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (s *ScriptIdioms) TableName() string {
	return "script_idioms"
}

func NewScriptIdiomsEntity(sIdioms *ScriptIdioms) *entity.ScriptIdioms {
	return &entity.ScriptIdioms{
		VideoID:   sIdioms.VideoID,
		ScriptID:  sIdioms.ScriptID,
		Idiom:     sIdioms.Word,
		CreatedAt: sIdioms.CreatedAt,
		UpdatedAt: sIdioms.UpdatedAt,
	}
}

func NewScriptIdiomsEntities(sIdioms []*ScriptIdioms, idiom []*Idioms) ([]*entity.ScriptIdioms, error) {
	return nil, nil
}
