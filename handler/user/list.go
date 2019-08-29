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
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	fmt.Printf("%v\n", r)
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
