package requests

import (
	"time"

	"github.com/HondaAo/snippet/src/pkg/services/quiz/entity"
)

type ListeningQuizRequest struct {
	VideoIDs []string `json:"video_ids"`
	Levels   []uint64 `json:"levels"`
	DateID   uint8    `json:"date"`
}

func (r *ListeningQuizRequest) GetDate() time.Time {
	switch r.DateID {
	case 1:
		return time.Now().AddDate(0, 0, -1)
	case 2:
		return time.Now().AddDate(0, 0, -7)
	case 3:
		return time.Now().AddDate(0, -1, 0)
	case 4:
		return time.Now().AddDate(-1, 0, 0)
	default:
		return time.Now().AddDate(-1, 0, 0)
	}
}

type CreateListeningQuiz struct {
	VideoID  string `json:"video_id"`
	ScriptID uint64 `json:"script_id"`
	Word     string `json:"word"`
	Level    uint8  `json:"level"`
}

func (r *CreateListeningQuiz) NewEntity() *entity.ListeningQuiz {
	return &entity.ListeningQuiz{
		VideoID:  r.VideoID,
		ScriptID: r.ScriptID,
		Word:     r.Word,
		Level:    r.Level,
	}
}
