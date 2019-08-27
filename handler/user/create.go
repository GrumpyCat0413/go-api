package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	. "go-api/handler"
	"go-api/pkg/errno"
)

func Create(c *gin.Context) {
	var r CreateRequest

	//参数绑定解析 post :Content-Type: application/json
	if err := c.Bind(&r); err != nil {
		//c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//展示了如何通过 Bind() 、 Param() 、 Query() 和 GetHeader() 来获取相应的参数
	/*
		   POST http://127.0.0.1:8080/v1/user/admin2?desc=hellodesc
		   Content-Type: application/json

		   {
			 "username":"asd",
			 "password":"111"
		   }

			c.Bind()  解析json
			c.Param() 解析 /:username 即admin2
			c.Query() 解析 ?号后面所带的参数
			c.GetHeader() 解析 header中 的指定字段的值
	*/

	admin2 := c.Param("username")
	log.Infof("URL username:%s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		//err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		//log.Errorf(err, "Get an error")
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
	}

	//if errno.IsErrUserNotFound(err) {
	//	log.Debug("err type is ErrUserNotFound")
	//}

	if r.Password == "" {
		//err = fmt.Errorf("password is empty")
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	//code, message := errno.DecodeErr(err)
	//c.JSON(http.StatusOK, gin.H{"code": code, "message": message})

	resp := CreateResponse{Username: r.Username}
	SendResponse(c, nil, resp)
}
