// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-20 16:37:39
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Water is the golang structure for table water.
type Water struct {
	Id        int64       `json:"id"        description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	WaterId   string      `json:"waterId"   description:""`
	WaterKey  string      `json:"waterKey"  description:""`
	Role      int         `json:"role"      description:""`
}
