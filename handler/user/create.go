package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "go-api/handler"
	"go-api/model"
	"go-api/pkg/errno"
	"go-api/util"
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

			c.Param() 返回url的参数值 解析 /:username 即admin2 路由中带有：的变量

			例如：GET /path?id=1234&name=Manu&value=
			c.Query() 读取url中的地址参数 解析 ?号后面所带的参数
			c.Query("id") == "1234"
			c.Query("name") == "Manu"
			c.Query("value") == ""
			c.Query("wtf") == ""
			c.GetHeader() 解析 header中 的指定字段的值 获取http的头部
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

func Create1(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	//参数绑定解析 post :Content-Type: application/json
	if err := c.Bind(&r); err != nil {
		//c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 验证 相关字段的有效性
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//用户 密码加密
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	//用户创建成功 返回用户名
	resp := CreateResponse{Username: r.Username}
	SendResponse(c, nil, resp)

}
