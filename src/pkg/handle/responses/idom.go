package responses

import (
	"github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
)

type IdiomResponse struct {
	Idiom     string `json:"idiom"`
	Meaning   string `json:"meaning"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewIdiomResponse(idiomEntity *entity.Idioms) IdiomResponse {
	level := getLevel(idiomEntity.Level)
	response := IdiomResponse{
		Idiom:     idiomEntity.Idiom,
		Meaning:   idiomEntity.Meaning,
		Level:     level,
		CreatedAt: idiomEntity.CreatedAt.Format("2006-01-02"),
		UpdatedAt: idiomEntity.UpdatedAt.Format("2006-01-02"),
	}

	return response
}

func getLevel(intLevel entity.IdiomLevel) string {
	switch intLevel {
	case 1:
		return "A1"
	case 2:
		return "A2"
	case 3:
		return "B1"
	case 4:
		return "B2"
	case 5:
		return "C1"
	default:
		return "F"
	}
}
