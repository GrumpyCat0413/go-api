package user

import (
	"github.com/gin-gonic/gin"
	"go-api/handler"
	"go-api/model"
	"go-api/pkg/errno"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
	}
	handler.SendResponse(c, nil, nil)
}
