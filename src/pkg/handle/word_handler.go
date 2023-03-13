package handle

import (
	"net/http"

	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/handle/responses"
	"github.com/HondaAo/snippet/src/pkg/services/word"
	"github.com/labstack/echo/v4"
)

type WordHandler interface {
	GetWords() echo.HandlerFunc
}

type wordHandler struct {
	wordUsecase word.WordUsecaseInterface
}

func NewWordHandler(usecase word.WordUsecaseInterface) *wordHandler {
	return &wordHandler{
		wordUsecase: usecase,
	}
}

func (w *wordHandler) GetWords() echo.HandlerFunc {
	return func(c echo.Context) error {
		newRequest := new(requests.WordRequest)
		if err := echo.QueryParamsBinder(c).Uint64s("levels", &newRequest.Levels).Uint64("amount", &newRequest.Amount).BindError(); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		words, err := w.wordUsecase.GetWords(newRequest.Levels, newRequest.Amount)
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		wordResponse := make([]*responses.WordResponse, 0, len(words))
		for _, word := range words {
			wordResponse = append(wordResponse, &responses.WordResponse{
				WordID:  word.WordID,
				Word:    word.Word,
				Mean:    word.Mean,
				Comment: word.Comment,
				Level:   word.Level,
			})
		}
		return c.JSON(http.StatusAccepted, words)
	}
}
