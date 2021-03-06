package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/seed95/clean-web-service/pkg/translate"
	server2 "github.com/seed95/clean-web-service/server"
	"strconv"
)

func getLang(c *gin.Context) []translate.Language {
	return server2.GetLanguage(c.GetHeader("Accept-Language"))
}

func getId(c *gin.Context) (uint, error) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
