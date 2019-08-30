package middleware

import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header 设置在返回包的 Header 中
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
