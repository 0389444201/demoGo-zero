// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userTableFieldNames          = builder.RawFieldNames(&UserTable{})
	userTableRows                = strings.Join(userTableFieldNames, ",")
	userTableRowsExpectAutoSet   = strings.Join(stringx.Remove(userTableFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userTableRowsWithPlaceHolder = strings.Join(stringx.Remove(userTableFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userTableModel interface {
		Insert(ctx context.Context, data *UserTable) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserTable, error)
		Update(ctx context.Context, data *UserTable) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTableModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserTable struct {
		Id       int64  `db:"id"`
		Email    string `db:"email"`
		Name     string `db:"name"`
		Gender   string `db:"gender"`
		Password string `db:"password"`
	}
)

func newUserTableModel(conn sqlx.SqlConn) *defaultUserTableModel {
	return &defaultUserTableModel{
		conn:  conn,
		table: "`user_table`",
	}
}

func (m *defaultUserTableModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserTableModel) FindOne(ctx context.Context, id int64) (*UserTable, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userTableRows, m.table)
	var resp UserTable
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

func (m *defaultUserTableModel) Insert(ctx context.Context, data *UserTable) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userTableRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Email, data.Name, data.Gender, data.Password)
	return ret, err
}

func (m *defaultUserTableModel) Update(ctx context.Context, data *UserTable) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userTableRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Email, data.Name, data.Gender, data.Password, data.Id)
	return err
}

func (m *defaultUserTableModel) tableName() string {
	return m.table
}