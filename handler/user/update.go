package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"go-api/handler"
	"go-api/model"
	"go-api/pkg/errno"
	"go-api/util"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// 从url 参数中获取user id
	userId, _ := strconv.Atoi(c.Param("id"))

	//绑定user数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
