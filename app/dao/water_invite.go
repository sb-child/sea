// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sea/app/dao/internal"
)

// waterInviteDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type waterInviteDao struct {
	*internal.WaterInviteDao
}

var (
	// WaterInvite is globally public accessible object for table water_invite operations.
	WaterInvite = waterInviteDao{
		internal.NewWaterInviteDao(),
	}
)

// Fill with you ideas below.