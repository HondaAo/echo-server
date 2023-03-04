package models

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
)

type Script struct {
	VideoID   string    `gorm:"primaryKey;autoIncrement:false"`
	ScriptID  uint64    `gorm:"primaryKey;autoIncrement:false"`
	Text      string    `gorm:"text"`
	Japanese  string    `gorm:"japanese"`
	TimeStamp float64   `gorm:"time_stamp"`
	CreatedAt time.Time `gorm:"crated_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (s *Script) TableName() string {
	return "scripts"
}

func NewScriptEntity(scriptModel *Script) *entity.Script {
	script := &entity.Script{
		VideoID:   scriptModel.VideoID,
		ScriptID:  scriptModel.ScriptID,
		Text:      scriptModel.Text,
		Japanese:  scriptModel.Japanese,
		TimeStamp: scriptModel.TimeStamp,
		CreatedAt: scriptModel.CreatedAt,
		UpdatedAt: scriptModel.UpdatedAt,
	}

	return script
}
