package router

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(otelgin.Middleware("httpserver"), Version, Log, CopyRequestHeader)

	r.GET("/healthz", HealthzHandler)

	return r
}
