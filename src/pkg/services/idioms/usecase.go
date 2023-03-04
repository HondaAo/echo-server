package idioms

import (
	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
)

type IdiomsUsecaseInterface interface {
	Find(idiom string) (*entity.Idioms, error)
	CreateIdioms(idioms *requests.CreateIdiomRequest) error
}

type idiomsUsecase struct {
	idiomRepository IdiomRepository
}

func NewIdiomUsecase(idiomRepository IdiomRepository) *idiomsUsecase {
	return &idiomsUsecase{
		idiomRepository,
	}
}

func (i *idiomsUsecase) Find(idiom string) (*entity.Idioms, error) {
	idioms, err := i.idiomRepository.Find(idiom)
	if err != nil {
		return nil, err
	}

	return idioms, nil
}

func (i *idiomsUsecase) CreateIdioms(request *requests.CreateIdiomRequest) error {
	idioms, err := request.NewEntities()
	if err != nil {
		return nil
	}

	if err := i.idiomRepository.Store(idioms); err != nil {
		return err
	}

	return nil
}
