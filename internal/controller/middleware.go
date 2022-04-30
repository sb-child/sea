package controller

import (
	"net/http"
	"sea/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func HeaderMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	selfWaterKey, err := service.WaterKey.GetSelfKey(r.Context())
	if err != nil {
		r.Response.Status = http.StatusServiceUnavailable
	}
	selfWaterId := selfWaterKey.GetKeyID()
	r.Response.Header().Add("water-id", selfWaterId)
}
