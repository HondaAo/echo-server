package handle

import "github.com/labstack/echo/v4"

func RegisterIdiomRoute(idiomGroup *echo.Group, handler IdiomHandler) {
	idiomGroup.GET("/:idiom", handler.Find())
	idiomGroup.POST("/create", handler.CreateIdioms())
}
