// Code generated by goctl. DO NOT EDIT.

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userHistoryFieldNames          = builder.RawFieldNames(&UserHistory{})
	userHistoryRows                = strings.Join(userHistoryFieldNames, ",")
	userHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(userHistoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(userHistoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLooklookTravelUserHistoryIdPrefix = "cache:looklookTravel:userHistory:id:"
)

type (
	userHistoryModel interface {
		Insert(ctx context.Context, data *UserHistory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserHistory, error)
		Update(ctx context.Context, data *UserHistory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserHistoryModel struct {
		sqlc.CachedConn
		table string
	}

	UserHistory struct {
		Id         int64     `db:"id"`
		HistoryId  int64     `db:"history_id"`
		UserId     int64     `db:"user_id"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"`
		DeleteTime time.Time `db:"delete_time"`
	}
)

func newUserHistoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserHistoryModel {
	return &defaultUserHistoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_history`",
	}
}

func (m *defaultUserHistoryModel) Delete(ctx context.Context, id int64) error {
	looklookTravelUserHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelUserHistoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookTravelUserHistoryIdKey)
	return err
}

func (m *defaultUserHistoryModel) FindOne(ctx context.Context, id int64) (*UserHistory, error) {
	looklookTravelUserHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelUserHistoryIdPrefix, id)
	var resp UserHistory
	err := m.QueryRowCtx(ctx, &resp, looklookTravelUserHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userHistoryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserHistoryModel) Insert(ctx context.Context, data *UserHistory) (sql.Result, error) {
	looklookTravelUserHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelUserHistoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userHistoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.HistoryId, data.UserId, data.DelState, data.Version, data.DeleteTime)
	}, looklookTravelUserHistoryIdKey)
	return ret, err
}

func (m *defaultUserHistoryModel) Update(ctx context.Context, data *UserHistory) error {
	looklookTravelUserHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelUserHistoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userHistoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.HistoryId, data.UserId, data.DelState, data.Version, data.DeleteTime, data.Id)
	}, looklookTravelUserHistoryIdKey)
	return err
}

func (m *defaultUserHistoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLooklookTravelUserHistoryIdPrefix, primary)
}

func (m *defaultUserHistoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userHistoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserHistoryModel) tableName() string {
	return m.table
}