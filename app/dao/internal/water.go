// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// WaterDao is the manager for logic model data accessing and custom defined data operations functions management.
type WaterDao struct {
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns WaterColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// WaterColumns defines and stores column names for table water.
type WaterColumns struct {
	WaterId    string //
	Self       string //
	PublicKey  string //
	PrivateKey string //
	VerifyStep string //
	CreatedAt  string //
	UpdatedAt  string //
	Banned     string //
}

//  waterColumns holds the columns for table water.
var waterColumns = WaterColumns{
	WaterId:    "water_id",
	Self:       "self",
	PublicKey:  "public_key",
	PrivateKey: "private_key",
	VerifyStep: "verify_step",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Banned:     "banned",
}

// NewWaterDao creates and returns a new DAO object for table data access.
func NewWaterDao() *WaterDao {
	return &WaterDao{
		Group:   "default",
		Table:   "water",
		Columns: waterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WaterDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WaterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WaterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}