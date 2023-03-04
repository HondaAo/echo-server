package idioms

import "github.com/HondaAo/snippet/src/pkg/services/idioms/entity"

type IdiomRepository interface {
	Store(idiom []*entity.Idioms) error
	Find(idiom string) (*entity.Idioms, error)
	FindIdioms(words []string) ([]*entity.Idioms, error)
}
