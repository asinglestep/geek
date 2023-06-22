package router

import (
	"os"
	"time"

	"geek/httpserver/utils/logger"

	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	version := os.Getenv("VERSION")
	c.Writer.Header().Add("version", version)
}

func CopyRequestHeader(c *gin.Context) {
	for k, vs := range c.Request.Header {
		for _, v := range vs {
			c.Writer.Header().Add(k, v)
		}
	}
}

func Log(c *gin.Context) {
	s := time.Now()

	c.Next()

	logger.WithContext(c.Request.Context()).Infof("[Access] path: %v, ip: %v, status code: %v, cost time: %v",
		c.Request.RequestURI, c.ClientIP(), c.Writer.Status(), time.Since(s))
}
