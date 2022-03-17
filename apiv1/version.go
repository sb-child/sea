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
