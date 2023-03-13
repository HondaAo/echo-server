package handle

import "github.com/labstack/echo/v4"

func RegisterWordRoute(wordGroup *echo.Group, handler WordHandler) {
	wordGroup.GET("/", handler.GetWords())
}
