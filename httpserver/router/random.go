package router

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var httpRequestDurationSecond = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Subsystem: "httpserver",
	Name:      "http_request_duration_seconds",
	Buckets:   []float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.08, 0.1, 0.2, 0.3, 0.5, 1.0, 1.5, 2.0},
}, []string{"path"})

func init() {
	prometheus.MustRegister(httpRequestDurationSecond)
}

// RandomHandler godoc
// @Summary 延时接口
// @Schemes
// @Description 添加0-2秒随机延时
// @Tags random
// @Success 200 {string} ok
// @Router /random [get]
func RandomHandler(c *gin.Context) {
	timer := prometheus.NewTimer(httpRequestDurationSecond.WithLabelValues(c.Request.RequestURI))
	defer timer.ObserveDuration()

	t := rand.New(rand.NewSource(time.Now().UnixMicro())).Intn(2000)
	time.Sleep(time.Duration(t) * time.Millisecond)
	c.String(http.StatusOK, "ok")
}
