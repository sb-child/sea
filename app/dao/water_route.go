// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sea/app/dao/internal"
)

// waterRouteDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type waterRouteDao struct {
	*internal.WaterRouteDao
}

var (
	// WaterRoute is globally public accessible object for table water_route operations.
	WaterRoute = waterRouteDao{
		internal.NewWaterRouteDao(),
	}
)

// Fill with you ideas below.
