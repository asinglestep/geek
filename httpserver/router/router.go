package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(Version, Log, CopyRequestHeader)
	r.GET("/healthz", HealthzHandler)

	return r
}
