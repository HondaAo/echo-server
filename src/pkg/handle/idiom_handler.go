package handle

import (
	"net/http"

	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/handle/responses"
	"github.com/HondaAo/snippet/src/pkg/services/idioms"
	"github.com/labstack/echo/v4"
)

type IdiomHandler interface {
	Find() echo.HandlerFunc
	CreateIdioms() echo.HandlerFunc
}

type idiomHandler struct {
	idiomUsecase idioms.IdiomsUsecaseInterface
}

func NewIdiomHandler(idiomUsecase idioms.IdiomsUsecaseInterface) *idiomHandler {
	return &idiomHandler{
		idiomUsecase,
	}
}

func (h *idiomHandler) Find() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.Param("idiom")
		idiom, err := h.idiomUsecase.Find(word)
		if err != nil {
			return err
		}

		response := responses.NewIdiomResponse(idiom)

		return c.JSON(http.StatusAccepted, response)
	}
}

func (h *idiomHandler) CreateIdioms() echo.HandlerFunc {
	return func(c echo.Context) error {
		idiomRequest := &requests.CreateIdiomRequest{}
		if err := c.Bind(idiomRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := h.idiomUsecase.CreateIdioms(idiomRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}
