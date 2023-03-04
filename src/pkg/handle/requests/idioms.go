package requests

import (
	"strings"

	"github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
)

type CreateIdiomRequest struct {
	Idioms []IdiomsRequest `json:"idioms"`
}

type IdiomsRequest struct {
	Idiom   string `json:"idiom"`
	Meaning string `json:"meaning"`
	Level   uint64 `json:"level"`
}

func (r *CreateIdiomRequest) NewEntities() ([]*entity.Idioms, error) {
	entities := make([]*entity.Idioms, 0, len(r.Idioms))
	for _, i := range r.Idioms {
		trimedIdiom := strings.ReplaceAll(i.Idiom, " ", "")
		entities = append(entities, &entity.Idioms{
			TrimedIdiom: trimedIdiom,
			Idiom:       i.Idiom,
			Meaning:     i.Meaning,
			Level:       entity.IdiomLevel(i.Level),
		})
	}

	return entities, nil
}
