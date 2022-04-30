package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WaterApiTransferReq struct {
	g.Meta  `path:"/t" method:"post"`
	Session string `p:"session" v:"required"`
	Origin  string `p:"origin" v:"required"`
	Next    string `p:"next" v:"required"`
	Target  string `p:"target" v:"required"`
	Type    string `p:"type" v:"required"`
	Value   string `p:"value" v:"required"`
}
type WaterApiTransferRes struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}
type WaterApiControlReq struct {
	g.Meta  `path:"/c" method:"post"`
	Session string `p:"session" v:"required"`
	Origin  string `p:"origin" v:"required"`
	Target  string `p:"target" v:"required"`
	Type    string `p:"type" v:"required"`
	Value   string `p:"value" v:"required"`
}
type WaterApiControlRes struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}
