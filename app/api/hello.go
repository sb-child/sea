package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Write("hello world")
}
