package word

import "github.com/HondaAo/snippet/src/pkg/services/word/entity"

type WordRepository interface {
	GetWords(level []uint64, amount uint64) ([]*entity.Word, error)
}
