package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAccountLogModel = (*customUserAccountLogModel)(nil)

type (
	// UserAccountLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAccountLogModel.
	UserAccountLogModel interface {
		userAccountLogModel
	}

	customUserAccountLogModel struct {
		*defaultUserAccountLogModel
	}
)

// NewUserAccountLogModel returns a model for the database table.
func NewUserAccountLogModel(conn sqlx.SqlConn) UserAccountLogModel {
	return &customUserAccountLogModel{
		defaultUserAccountLogModel: newUserAccountLogModel(conn),
	}
}
