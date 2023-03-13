package word

import "github.com/HondaAo/snippet/src/pkg/services/word/entity"

type WordUsecaseInterface interface {
	GetWords(level []uint64, amount uint64) ([]*entity.Word, error)
}

type wordUsecase struct {
	wordReopository WordRepository
}

func NewWordUsecase(
	wordRepository WordRepository,
) *wordUsecase {
	return &wordUsecase{
		wordReopository: wordRepository,
	}
}

func (w *wordUsecase) GetWords(level []uint64, amount uint64) ([]*entity.Word, error) {
	words, err := w.wordReopository.GetWords(level, amount)
	if err != nil {
		return nil, err
	}

	return words, nil
}
