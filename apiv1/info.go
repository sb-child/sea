package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type VersionReq struct {
	g.Meta `path:"/ver" method:"get"`
}
type VersionRes struct {
	g.Meta    `mime:"application/json"`
	BuildInfo map[string]string `json:"info"`
}

type PublicKeyReq struct {
	g.Meta `path:"/publickey" method:"get"`
}
type PublicKeyRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Id     string `json:"id"`
}
