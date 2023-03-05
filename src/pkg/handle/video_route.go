package handle

import "github.com/labstack/echo/v4"

func RegisterVideoRoute(videoGroup *echo.Group, handler Handler) {
	videoGroup.GET("/videos", handler.GetVideos())
	videoGroup.POST("/create", handler.StoreVideo())
	videoGroup.PUT("/update", handler.UpdateVideo())
	videoGroup.GET("/:video_id", handler.GetVideo())
	videoGroup.PUT("/status/:video_id", handler.ChangeStatus())
}
