package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/seed95/clean-web-service/internal/server/rest"
	"github.com/seed95/clean-web-service/pkg/translate"
	"strconv"
)

func getLang(c echo.Context) []translate.Language {
	return rest.GetLanguage(c.Request().Header.Get("Accept-Language"))
}

func getId(c echo.Context) (uint, error) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
