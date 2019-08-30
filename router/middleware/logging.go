package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"go-api/handler"
	"go-api/pkg/errno"
	"io/ioutil"
	"regexp"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path
		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		if path == "sd/helth" || path == "/sd/ram" || path == "/sd/cpu" || path == "sd/disk" {
			return
		}

		/*
			该中间件需要截获 HTTP 的请求信息，然后打印请求信息，
			因为 HTTP 的请求 Body，在 读取过后会被置空，所以这里读取完后会重新赋值：
		*/
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		//Restore
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start) //等待时间

		code, message := -1, ""

		var response handler.Response

		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}
		// 记录 耗时、请求ip、http方法http路径 返回的code和message
		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
