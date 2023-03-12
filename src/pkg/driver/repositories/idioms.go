package repositories

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"github.com/HondaAo/snippet/src/pkg/services/idioms/entity"

	"gorm.io/gorm"
)

type idiomsRepository struct {
	db *gorm.DB
}

func NewIdiomsRepository(db *gorm.DB) *idiomsRepository {
	return &idiomsRepository{
		db: db,
	}
}

func (r *idiomsRepository) Find(idiom string) (*entity.Idioms, error) {
	var idioms = &models.Idioms{}

	if result := r.db.First(idioms, "trimed_idiom = ?", idiom); result.Error != nil {
		return nil, result.Error
	}

	return models.NewIdiomEntity(idioms), nil
}

func (r *idiomsRepository) FindIdioms(words []string) ([]*entity.Idioms, error) {
	idioms := []*models.Idioms{}
	if result := r.db.Where("trimed_idiom IN ?", words).Find(&idioms); result.Error != nil {
		return nil, result.Error
	}

	idiomEntities := make([]*entity.Idioms, 0, len(idioms))
	for _, i := range idioms {
		idiomEntities = append(idiomEntities, models.NewIdiomEntity(i))
	}

	return idiomEntities, nil
}

func (r *idiomsRepository) Store(idioms []*entity.Idioms) error {
	idiomModels := make([]*models.Idioms, 0, len(idioms))
	for _, idiom := range idioms {
		idiomModels = append(idiomModels, &models.Idioms{
			TrimedIdiom: idiom.TrimedIdiom,
			Idiom:       idiom.Idiom,
			Meaning:     idiom.Meaning,
			Level:       uint64(idiom.Level),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	if result := r.db.Create(idiomModels); result.Error != nil {
		return result.Error
	}

	return nil
}
