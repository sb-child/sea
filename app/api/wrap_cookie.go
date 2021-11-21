package api

import (
	"net/http"
	serviceWater "sea/app/service/water"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CookieMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
	selfWaterKey, err := serviceWater.WaterKey.GetSelfKey(r.Context())
	if err != nil {
		r.Response.Status = http.StatusServiceUnavailable
	}
	selfWaterId := selfWaterKey.GetKeyID()
	r.Cookie.Set("water-id", selfWaterId)
}
