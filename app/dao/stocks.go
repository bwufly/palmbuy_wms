// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wms/app/dao/internal"
)

// stocksDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type stocksDao struct {
	*internal.StocksDao
}

var (
	// Stocks is globally public accessible object for table stocks operations.
	Stocks = &stocksDao{
		internal.Stocks,
	}
)

// Fill with you ideas below.