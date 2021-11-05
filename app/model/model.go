// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Water is the golang structure for table water.
type Water struct {
	WaterId    string      `orm:"water_id"    json:"waterId"    description:""`
	Self       bool        `orm:"self"        json:"self"       description:""`
	Key        string      `orm:"key"         json:"key"        description:""`
	VerifyStep int         `orm:"verify_step" json:"verifyStep" description:""`
	CreatedAt  *gtime.Time `orm:"created_at"  json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updatedAt"  description:""`
	Banned     bool        `orm:"banned"      json:"banned"     description:""`
	Url        string      `orm:"url"         json:"url"        description:""`
	Session    string      `orm:"session"     json:"session"    description:""`
}

// WaterPing is the golang structure for table water_ping.
type WaterPing struct {
	WaterFrom string `orm:"water_from" json:"waterFrom" description:""`
	WaterTo   string `orm:"water_to"   json:"waterTo"   description:""`
	Delay     int64  `orm:"delay"      json:"delay"     description:""`
}

// WaterRoute is the golang structure for table water_route.
type WaterRoute struct {
	WaterTo  string `orm:"water_to" json:"waterTo"  description:""`
	Route    string `orm:"route"    json:"route"    description:""`
	Delay    int64  `orm:"delay"    json:"delay"    description:""`
	Disabled bool   `orm:"disabled" json:"disabled" description:""`
	Skipped  bool   `orm:"skipped"  json:"skipped"  description:""`
}
