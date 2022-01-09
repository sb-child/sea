// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Water is the golang structure of table water for DAO operations like Where/Data.
type Water struct {
	g.Meta        `orm:"table:water, dto:true"`
	WaterId       interface{} //
	VerifySession interface{} //
	Key           interface{} //
	Url           interface{} //
	IsBanned      interface{} //
	IsVerified    interface{} //
	IsReviewed    interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	IsSelf        interface{} //
}
