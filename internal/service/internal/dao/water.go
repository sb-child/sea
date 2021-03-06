// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sea/internal/service/internal/dao/internal"
)

// waterDao is the data access object for table water.
// You can define custom methods on it to extend its functionality as you wish.
type waterDao struct {
	*internal.WaterDao
}

var (
	// Water is globally public accessible object for table water operations.
	Water = waterDao{
		internal.NewWaterDao(),
	}
)

// Fill with you ideas below.
