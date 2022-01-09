package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// type HelloReq struct {
// 	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
// }
// type HelloRes struct {
// 	g.Meta `mime:"text/html" example:"string"`
// }

type VersionReq struct {
	g.Meta `path:"/ver" method:"get"`
}
type VersionRes struct {
	g.Meta `mime:"application/json"`
}

type WaterApiInviteStep1Req struct {
	g.Meta `path:"/auth/join/1" method:"post"`
}
type WaterApiInviteStep1Res struct {
	g.Meta `mime:"application/json"`
}

type WaterApiInviteStep2Req struct {
	g.Meta `path:"/auth/join/2" method:"post"`
}
type WaterApiInviteStep2Res struct {
	g.Meta `mime:"application/json"`
}

type AdminWaterAddReq struct {
	g.Meta `path:"/water/add" method:"post"`
}
type AdminWaterAddRes struct {
	g.Meta `mime:"application/json"`
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
