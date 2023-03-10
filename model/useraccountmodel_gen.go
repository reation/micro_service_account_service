// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAccountFieldNames          = builder.RawFieldNames(&UserAccount{})
	userAccountRows                = strings.Join(userAccountFieldNames, ",")
	userAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(userAccountFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(userAccountFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	STATE_NORMAL  int64 = 1
	STATE_UNUSUAL int64 = 21
	STATE_FORZEN  int64 = 31

	TYPE_PAY_ORDER    int64 = 1
	TYPE_RETURN_ORDER int64 = 2
	TYPE_COMPANSATE   int64 = 3
	TYPE_PAY_MEMBER   int64 = 4
)

type (
	userAccountModel interface {
		Insert(ctx context.Context, data *UserAccount) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAccount, error)
		Update(ctx context.Context, data *UserAccount) error
		Delete(ctx context.Context, id int64) error
		GetAccountInfoByUserID(ctx context.Context, userID int64) (*UserAccount, error)
	}

	defaultUserAccountModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserAccount struct {
		Id          int64     `db:"id"`
		UserId      int64     `db:"user_id"`      // 用户ID
		PayPassword string    `db:"pay_password"` // 余额支付密码的MD5
		Balance     float64   `db:"balance"`      // 账户余额
		State       int64     `db:"state"`        // 账户状态 1：正常， 21:异常， 31：冻结
		Code        int64     `db:"code"`         // 修改码
		UpdateTime  time.Time `db:"update_time"`
		CreateTime  time.Time `db:"create_time"`
	}
)

func newUserAccountModel(conn sqlx.SqlConn) *defaultUserAccountModel {
	return &defaultUserAccountModel{
		conn:  conn,
		table: "`user_account`",
	}
}

func (m *defaultUserAccountModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserAccountModel) FindOne(ctx context.Context, id int64) (*UserAccount, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAccountRows, m.table)
	var resp UserAccount
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) GetAccountInfoByUserID(ctx context.Context, userID int64) (*UserAccount, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userAccountRows, m.table)
	var resp UserAccount
	err := m.conn.QueryRowCtx(ctx, &resp, query, userID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAccountModel) Insert(ctx context.Context, data *UserAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userAccountRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.PayPassword, data.Balance, data.State, data.Code)
	return ret, err
}

func (m *defaultUserAccountModel) Update(ctx context.Context, data *UserAccount) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAccountRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.PayPassword, data.Balance, data.State, data.Code, data.Id)
	return err
}

func (m *defaultUserAccountModel) tableName() string {
	return m.table
}
