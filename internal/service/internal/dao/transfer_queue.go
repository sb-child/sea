// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"sea/internal/service/internal/dao/internal"
)

// transferQueueDao is the data access object for table transfer_queue.
// You can define custom methods on it to extend its functionality as you wish.
type transferQueueDao struct {
	*internal.TransferQueueDao
}

var (
	// TransferQueue is globally public accessible object for table transfer_queue operations.
	TransferQueue = transferQueueDao{
		internal.NewTransferQueueDao(),
	}
)

// Fill with you ideas below.