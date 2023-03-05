package handle

import (
	"net/http"
	"strconv"

	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/handle/responses"
	"github.com/HondaAo/snippet/src/pkg/services/videos"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetVideo() echo.HandlerFunc
	GetVideos() echo.HandlerFunc
	StoreVideo() echo.HandlerFunc
	UpdateVideo() echo.HandlerFunc
	ChangeStatus() echo.HandlerFunc
}

type videoHandler struct {
	useCase videos.VideoUsecaseInterface
}

func NewVideoHandler(useCase videos.VideoUsecaseInterface) *videoHandler {
	return &videoHandler{
		useCase: useCase,
	}
}

func (v *videoHandler) GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("video_id")
		video, scripts, scriptIdioms, idioms, err := v.useCase.GetVideoByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		videoResponse := responses.NewResponse(video, scripts, scriptIdioms, idioms)

		return c.JSON(http.StatusAccepted, videoResponse)
	}
}

func (v *videoHandler) GetVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		level, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return err
		}
		categoryID, err := strconv.Atoi(c.QueryParam("category_id"))
		if err != nil {
			return err
		}
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return err
		}

		videos, err := v.useCase.GetVideos(uint64(limit), uint64(level), uint64(categoryID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		return c.JSON(http.StatusAccepted, videos)
	}
}

func (v *videoHandler) StoreVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		videoRequest := &requests.VideoRequest{}
		if err := c.Bind(videoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := v.useCase.CreateVideo(videoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}

func (v *videoHandler) UpdateVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		videoRequest := &requests.VideoUpdateRequest{}
		if err := c.Bind(videoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := v.useCase.UpdateVideo(videoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}

func (v *videoHandler) ChangeStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("video_id")
		if err := v.useCase.ChangeDisplayStatus(id); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}
