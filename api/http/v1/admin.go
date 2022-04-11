package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminWaterAddReq struct {
	g.Meta `path:"/water/add" method:"post"`
	URL    string `p:"url" v:"required"`
}
type AdminWaterAddRes struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}

type AdminWaterDeleteReq struct {
	g.Meta `path:"/water/delete" method:"post"`
}
type AdminWaterDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type AdminWaterEditReq struct {
	g.Meta `path:"/water/edit" method:"post"`
}
type AdminWaterEditRes struct {
	g.Meta `mime:"application/json"`
}

type AdminWaterQueryReq struct {
	g.Meta `path:"/water/query" method:"post"`
}
type AdminWaterQueryRes struct {
	g.Meta `mime:"application/json"`
}
