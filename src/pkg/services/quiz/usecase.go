package quiz

import (
	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/services/quiz/entity"
)

type QuizUsecaseInterface interface {
	GetListeningQuizs() ([]*entity.ListeningQuiz, error)
}

type quizUsecase struct {
	quizRepository QuizRepository
}

func NewQuizUsecase(
	quizRepository QuizRepository,
) *quizUsecase {
	return &quizUsecase{
		quizRepository: quizRepository,
	}
}

func (q *quizUsecase) GetListeningQuizs(request requests.ListeningQuizRequest) ([]*entity.ListeningQuiz, error) {
	quizs, err := q.quizRepository.GetListringQuizs(request.VideoIDs, request.Levels, request.GetDate())
	if err != nil {
		return nil, err
	}

	return quizs, nil
}

func (q *quizUsecase) Store(request requests.CreateListeningQuiz) error {
	if err := q.quizRepository.Store(request.NewEntity()); err != nil {
		return err
	}

	return nil
}
