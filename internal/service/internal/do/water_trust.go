// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-20 16:37:39
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WaterTrust is the golang structure of table water_trust for DAO operations like Where/Data.
type WaterTrust struct {
	g.Meta    `orm:"table:water_trust, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	WaterId   interface{} //
	IsBanned  interface{} //
}
