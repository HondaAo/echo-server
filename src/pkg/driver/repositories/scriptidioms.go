package repositories

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
	"gorm.io/gorm"
)

type scriptIdiomsRepository struct {
	db *gorm.DB
}

func NewScriptIdiomsRepository(db *gorm.DB) *scriptIdiomsRepository {
	return &scriptIdiomsRepository{
		db: db,
	}
}

func (r *scriptIdiomsRepository) Find(words []string) ([]*entity.ScriptIdioms, error) {
	var scriptIdioms []*models.ScriptIdioms
	if result := r.db.Where("word IN ?", words).Find(&scriptIdioms); result.Error != nil {
		return nil, result.Error
	}

	scriptEntities := make([]*entity.ScriptIdioms, 0, len(scriptIdioms))
	for _, s := range scriptIdioms {
		scriptEntities = append(scriptEntities, models.NewScriptIdiomsEntity(s))
	}
	return scriptEntities, nil
}

func (r *scriptIdiomsRepository) FindByVideoID(videoId string) ([]*entity.ScriptIdioms, error) {
	var scriptIdioms []*models.ScriptIdioms
	if result := r.db.Where("video_id = ?", videoId).Find(&scriptIdioms); result.Error != nil {
		return nil, result.Error
	}
	scriptIdiomsEntities := make([]*entity.ScriptIdioms, 0, len(scriptIdioms))
	for _, s := range scriptIdioms {
		scriptIdiomsEntities = append(scriptIdiomsEntities, models.NewScriptIdiomsEntity(s))
	}
	return scriptIdiomsEntities, nil
}

func (r *scriptIdiomsRepository) Store(words []entity.ScriptIdioms) error {
	scriptIdiomsModels := make([]*models.ScriptIdioms, 0, len(words))
	for _, s := range words {
		scriptIdiomsModels = append(scriptIdiomsModels, &models.ScriptIdioms{
			VideoID:   s.VideoID,
			ScriptID:  s.ScriptID,
			Word:      s.Idiom,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	if result := r.db.Create(scriptIdiomsModels); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *scriptIdiomsRepository) Delete(videoID string) error {
	if result := r.db.Where("video_id = ?", videoID).Delete(&models.ScriptIdioms{}); result.Error != nil {
		return result.Error
	}

	return nil
}
