package quiz

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/quiz/entity"
)

type QuizRepository interface {
	GetListringQuizs(videoID []string, level []uint64, date time.Time) ([]*entity.ListeningQuiz, error)
	Store(entity *entity.ListeningQuiz) error
}
