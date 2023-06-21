package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	swag "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))

	r.Use(otelgin.Middleware("httpserver"), Version, Log, CopyRequestHeader)

	r.GET("/healthz", HealthzHandler)
	r.GET("/random", RandomHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
