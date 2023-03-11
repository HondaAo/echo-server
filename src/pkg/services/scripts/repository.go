package scripts

import "github.com/HondaAo/snippet/src/pkg/services/scripts/entity"

type ScriptRepository interface {
	FindByVideoID(videoID string) ([]*entity.Script, error)
	Store(videoID string, scripts []entity.Script) error
	Update(script *entity.Script) error
	Delete(videoID string) error
}
