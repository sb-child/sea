package apiv1

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gbuild"
)

var GetVersion = getVersionApi{}

type getVersionApi struct{}

type getVersionJson struct {
	BuildInfo map[string]string `json:"info"`
}

func (*getVersionApi) BuildInfo(r *ghttp.Request) {
	r.Response.WriteJsonExit(&getVersionJson{
		BuildInfo: gbuild.Info(),
	})
}
