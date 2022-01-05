package utils

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func ParseError(r *ghttp.Request, err error) {
	r.Response.WriteStatus(500)
	r.Response.WriteJsonExit(map[string]interface{}{
		"code":    500,
		"message": err.Error(),
	})
}
