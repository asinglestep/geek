package router

import (
	"os"

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
	c.Next()

	logger.WithContext(c.Request.Context()).Infof("[Access] path: %v, ip: %v, status code: %v",
		c.Request.RequestURI, c.ClientIP(), c.Writer.Status())
}
