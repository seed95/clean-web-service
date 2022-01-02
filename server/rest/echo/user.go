package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/log"
	"github.com/seed95/clean-web-service/server/params"
	"net/http"
)

func (h *server) createUser(c echo.Context) error {

	lang := getLang(c)

	reqParam := new(params.CreateUserRequest)
	if err := c.Bind(reqParam); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "createUser",
			Message:  h.translator.Translate(err.Error()),
		})
		return echo.NewHTTPError(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
	}

	user, err := params.ConvertRequest(reqParam)
	if err != nil {
		msg, code := derrors.HttpError(err)

		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "createUser",
			Params:   map[string]interface{}{"user": reqParam},
			Message:  h.translator.Translate(msg, lang...),
		})

		return echo.NewHTTPError(code, h.translator.Translate(msg, lang...))
	}

	user, err = h.userService.CreateUser(user)
	if err != nil {
		msg, code := derrors.HttpError(err)
		return echo.NewHTTPError(code, h.translator.Translate(msg, lang...))
	}

	return c.JSON(http.StatusOK, params.ConvertCreateUserResponse(user))
}

func (h *server) getUser(c echo.Context) error {

	lang := getLang(c)

	id, err := getId(c)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "getUser",
			Message:  err.Error(),
		})

		return echo.NewHTTPError(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
	}

	user, err := h.userService.GetUserById(id)
	if err != nil {
		msg, code := derrors.HttpError(err)
		return echo.NewHTTPError(code, h.translator.Translate(msg, lang...))
	}

	return c.JSON(http.StatusOK, params.ConvertGetUserResponse(user))

}

func (h *server) deleteUser(c echo.Context) error {

	lang := getLang(c)

	id, err := getId(c)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "getUser",
			Message:  err.Error(),
		})

		return echo.NewHTTPError(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
	}

	user, err := h.userService.DeleteUserById(id)
	if err != nil {
		msg, code := derrors.HttpError(err)
		return echo.NewHTTPError(code, h.translator.Translate(msg, lang...))
	}

	return c.JSON(http.StatusOK, params.ConvertDeleteUserResponse(user))

}
