// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sea/internal/service/internal/dao/internal"
)

// usersDao is the data access object for table users.
// You can define custom methods on it to extend its functionality as you wish.
type usersDao struct {
	*internal.UsersDao
}

var (
	// Users is globally public accessible object for table users operations.
	Users = usersDao{
		internal.NewUsersDao(),
	}
)

// Fill with you ideas below.
