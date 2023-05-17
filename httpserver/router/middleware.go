package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

	glog.Infof("[Access] ip: %v, status code: %v", c.ClientIP(), c.Writer.Status())
}
