// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WaterRoute is the golang structure of table water_route for DAO operations like Where/Data.
type WaterRoute struct {
	g.Meta   `orm:"table:water_route, dto:true"`
	WaterTo  interface{} //
	Route    interface{} //
	Delay    interface{} //
	Disabled interface{} //
	Skipped  interface{} //
}