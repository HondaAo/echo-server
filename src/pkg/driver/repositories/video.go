package repositories

import (
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

	return models.NewEntity(video)
}

func (r *videoRepository) Store(video *entity.Video) error {
	videoModel := &models.Video{
		VideoID:   video.VideoID,
		Title:     video.Title,
		URL:       video.URL,
		Start:     video.Start,
		End:       video.End,
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

func (r *videoRepository) Delete(videoID string) error {
	return nil
}
