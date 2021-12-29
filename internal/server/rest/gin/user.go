package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/internal/server/params"
	"github.com/seed95/clean-web-service/pkg/derrors"
	"github.com/seed95/clean-web-service/pkg/log"
	"net/http"
)

func (h *server) createUser(c *gin.Context) {

	lang := getLang(c)

	reqParam := new(params.CreateUserRequest)
	if err := c.Bind(reqParam); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "createUser",
			Message:  h.translator.Translate(err.Error()),
		})

		c.JSON(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
		return
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

		c.JSON(code, h.translator.Translate(msg, lang...))
		return
	}

	user, err = h.userService.CreateUser(user)
	if err != nil {
		msg, code := derrors.HttpError(err)
		c.JSON(code, h.translator.Translate(msg, lang...))
		return
	}

	c.JSON(http.StatusOK, params.ConvertCreateUserResponse(user))
	
}

func (h *server) getUser(c *gin.Context) {

	lang := getLang(c)

	id, err := getId(c)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "getUser",
			Message:  err.Error(),
		})

		c.JSON(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
		return
	}

	user, err := h.userService.GetUserById(id)
	if err != nil {
		msg, code := derrors.HttpError(err)

		c.JSON(code, h.translator.Translate(msg, lang...))
		return
	}

	c.JSON(http.StatusOK, params.ConvertGetUserResponse(user))

}

func (h *server) deleteUser(c *gin.Context) {

	lang := getLang(c)

	id, err := getId(c)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.user",
			Function: "getUser",
			Message:  err.Error(),
		})

		c.JSON(http.StatusBadRequest, h.translator.Translate(messages.ParseQueryError, lang...))
		return
	}

	user, err := h.userService.DeleteUserById(id)
	if err != nil {
		msg, code := derrors.HttpError(err)
		c.JSON(code, h.translator.Translate(msg, lang...))
		return
	}

	c.JSON(http.StatusOK, params.ConvertDeleteUserResponse(user))

}
