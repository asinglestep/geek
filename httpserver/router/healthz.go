package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthzHandler godoc
// @Summary 健康监测接口
// @Schemes
// @Description 健康监测
// @Tags healthz
// @Success 200 {string} ok
// @Router /healthz [get]
func HealthzHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
