package repositories

import (
	"log"
	"time"

	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
	"gorm.io/gorm"
)

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *videoRepository {
	return &videoRepository{
		db: db,
	}
}

func (r *videoRepository) Find(videoID string) (*entity.Video, error) {
	video := &models.Video{}
	r.db.First(video, "video_id = ?", videoID)

	return models.NewEntity(video), nil
}

func (r *videoRepository) Store(video *entity.Video) error {
	videoModel := &models.Video{
		VideoID:   video.VideoID,
		Title:     video.Title,
		URL:       video.URL,
		Start:     video.Start,
		End:       video.End,
		Length:    video.End - video.Start,
		Level:     uint64(video.Level),
		Display:   video.Display,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := r.db.Create(videoModel)
	if result.Error != nil {
		// result.Error.Error()
		return result.Error
	}

	return nil
}

func (r *videoRepository) Update(video *entity.Video) error {
	videoModel := &models.Video{
		VideoID:   video.VideoID,
		Title:     video.Title,
		URL:       video.URL,
		Start:     video.Start,
		End:       video.End,
		Level:     uint64(video.Level),
		Display:   video.Display,
		UpdatedAt: time.Now(),
	}
	if result := r.db.Where("video_id = ?", videoModel.VideoID).Updates(&videoModel); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *videoRepository) ChangeStatus(videoID string) error {
	if result := r.db.Model(&models.Video{}).Where("video_id = ?", videoID).Update("display", true); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *videoRepository) FindMany(condition entity.SearchCondition) ([]*entity.Video, error) {
	log.Println(condition)
	videos := []*models.Video{}
	result := r.db.Limit(int(condition.Limit)).Where("created_at > ?", condition.Date)
	if len(condition.VideoIDs) > 0 {
		result.Where("video_id IN ?", condition.VideoIDs)
	}
	if len(condition.CategoryIDs) > 0 {
		result.Where("category_id IN ?", condition.CategoryIDs)
	}
	if len(condition.Levels) != 0 {
		result.Where("level IN ?", condition.Levels)
	}
	if len(condition.TypeIDs) > 0 {
		result.Where("type IN ?", condition.TypeIDs)
	}
	if len(condition.Length) > 0 {
		result.Where("length BETWEEN ? AND ?", condition.Length[0], condition.Length[1])
	}
	result.Find(&videos)

	if result.Error != nil {
		return nil, result.Error
	}

	entities := make([]*entity.Video, 0, len(videos))
	for _, v := range videos {
		entities = append(entities, models.NewEntity(v))
	}

	return entities, nil
}
func (r *videoRepository) Delete(videoID string) error {
	if result := r.db.Where("video_id = ?", videoID).Delete(&models.Video{}); result.Error != nil {
		return result.Error
	}
	return nil
}
