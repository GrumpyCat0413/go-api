package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/handler"
	"go-api/pkg/errno"
	"go-api/service"
)

// List list the users in the database.
func List(c *gin.Context) {
	fmt.Println(c.Request.Method, c.ContentType())

	var r ListRequest

	// 对于get 带有json body的请求 应该是用下面的bind方法
	if err := c.ShouldBindJSON(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err = json.Unmarshal(bytes, &r); err != nil {
	//	handler.SendResponse(c, err, nil)
	//	return
	//}

	fmt.Println(r)

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
