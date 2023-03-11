package scriptidioms

import "github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"

type ScriptIdiomsRepository interface {
	Find(words []string) ([]*entity.ScriptIdioms, error)
	FindByVideoID(videoID string) ([]*entity.ScriptIdioms, error)
	Store(word []entity.ScriptIdioms) error
	Delete(videoID string) error
}
