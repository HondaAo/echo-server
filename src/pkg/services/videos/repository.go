package videos

import "github.com/HondaAo/snippet/src/pkg/services/videos/entity"

type VideoRepository interface {
	Find(videoID string) (*entity.Video, error)
	ChangeStatus(videoID string) error
	Store(video *entity.Video) error
	Update(video *entity.Video) error
	Delete(videoID string) error
}
