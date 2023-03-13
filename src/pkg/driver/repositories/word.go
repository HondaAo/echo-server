package repositories

import (
	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"github.com/HondaAo/snippet/src/pkg/services/word/entity"
	"gorm.io/gorm"
)

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) *wordRepository {
	return &wordRepository{
		db: db,
	}
}

func (r *wordRepository) GetWords(level []uint64, amount uint64) ([]*entity.Word, error) {
	words := []*models.Words{}
	if result := r.db.Limit(int(amount)).Where("level IN ?", level).Order("rand()").Find(&words); result.Error != nil {
		return nil, result.Error
	}

	wordEntities := make([]*entity.Word, 0, len(words))
	for _, word := range words {
		wordEntities = append(wordEntities, &entity.Word{
			WordID:    word.WordID,
			Word:      word.Word,
			Mean:      word.Mean,
			Comment:   word.Comment,
			CreatedAt: word.CreatedAt,
			UpdatedAt: word.UpdatedAt,
		})
	}

	return wordEntities, nil
}
