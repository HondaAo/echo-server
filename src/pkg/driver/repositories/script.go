package repositories

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
	"gorm.io/gorm"
)

type scriptRepository struct {
	db *gorm.DB
}

func NewScriptRepository(db *gorm.DB) *scriptRepository {
	return &scriptRepository{
		db: db,
	}
}

func (s *scriptRepository) FindByVideoID(videoID string) ([]*entity.Script, error) {
	var sciprts []*models.Script
	result := s.db.Where("video_id = ?", videoID).Find(&sciprts)
	if result.Error != nil {
		return nil, result.Error
	}
	scriptEntities := make([]*entity.Script, 0, len(sciprts))
	for _, s := range sciprts {
		scriptEntities = append(scriptEntities, models.NewScriptEntity(s))
	}
	return scriptEntities, nil
}

func (s *scriptRepository) Store(videoID string, scripts []entity.Script) error {
	scriptModels := make([]*models.Script, 0, len(scripts))
	for _, s := range scripts {
		scriptModels = append(scriptModels, &models.Script{
			VideoID:   videoID,
			ScriptID:  s.ScriptID,
			Text:      s.Text,
			Japanese:  s.Japanese,
			TimeStamp: s.TimeStamp,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	result := s.db.Create(scriptModels)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *scriptRepository) Update(script *entity.Script) error {
	scriptModel := &models.Script{
		VideoID:   script.VideoID,
		ScriptID:  script.ScriptID,
		TimeStamp: script.TimeStamp,
		Japanese:  script.Japanese,
		Text:      script.Text,
		UpdatedAt: time.Now(),
	}

	result := s.db.Where("video_id = ?", scriptModel.VideoID).Where("script_id", scriptModel.ScriptID).Updates(&scriptModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *scriptRepository) Delete(videoID string) error {
	if result := s.db.Where("video_id = ?", videoID).Delete(&models.Script{}); result.Error != nil {
		return result.Error
	}

	return nil
}
