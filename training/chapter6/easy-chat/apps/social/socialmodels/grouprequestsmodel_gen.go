// Code generated by goctl. DO NOT EDIT.

package socialmodels

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupRequestsFieldNames          = builder.RawFieldNames(&GroupRequests{})
	groupRequestsRows                = strings.Join(groupRequestsFieldNames, ",")
	groupRequestsRowsExpectAutoSet   = strings.Join(stringx.Remove(groupRequestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupRequestsRowsWithPlaceHolder = strings.Join(stringx.Remove(groupRequestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGroupRequestsIdPrefix = "cache:groupRequests:id:"
)

type (
	groupRequestsModel interface {
		Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		Insert(ctx context.Context, data *GroupRequests) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*GroupRequests, error)
		FindByGroupIdAndReqId(ctx context.Context, groupId, reqId string) (*GroupRequests, error)
		ListNoHandler(ctx context.Context, groupId string) ([]*GroupRequests, error)
		Update(ctx context.Context, session sqlx.Session, data *GroupRequests) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultGroupRequestsModel struct {
		sqlc.CachedConn
		table string
	}

	GroupRequests struct {
		Id            uint64         `db:"id"`
		ReqId         string         `db:"req_id"`
		GroupId       string         `db:"group_id"`
		ReqMsg        sql.NullString `db:"req_msg"`
		ReqTime       sql.NullTime   `db:"req_time"`
		JoinSource    sql.NullInt64  `db:"join_source"`
		InviterUserId sql.NullString `db:"inviter_user_id"`
		HandleUserId  sql.NullString `db:"handle_user_id"`
		HandleTime    sql.NullTime   `db:"handle_time"`
		HandleResult  sql.NullInt64  `db:"handle_result"`
	}
)

func newGroupRequestsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupRequestsModel {
	return &defaultGroupRequestsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`group_requests`",
	}
}

func (m *defaultGroupRequestsModel) Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultGroupRequestsModel) Delete(ctx context.Context, id uint64) error {
	groupRequestsIdKey := fmt.Sprintf("%s%v", cacheGroupRequestsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, groupRequestsIdKey)
	return err
}

func (m *defaultGroupRequestsModel) FindOne(ctx context.Context, id uint64) (*GroupRequests, error) {
	groupRequestsIdKey := fmt.Sprintf("%s%v", cacheGroupRequestsIdPrefix, id)
	var resp GroupRequests
	err := m.QueryRowCtx(ctx, &resp, groupRequestsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRequestsRows, m.table)
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

func (m *defaultGroupRequestsModel) FindByGroupIdAndReqId(ctx context.Context, groupId, reqId string) (*GroupRequests, error) {
	query := fmt.Sprintf("query %s from %s where `req_id` = ? and `group_id` = ?",
		groupRequestsRows, m.table)

	var resp GroupRequests
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, reqId, groupId)
	return &resp, err
}

func (m *defaultGroupRequestsModel) ListNoHandler(ctx context.Context, groupId string) ([]*GroupRequests, error) {
	query := fmt.Sprintf("select %s from %s where `group_id` = ? and `handle_result` = 1",
		groupRequestsRows, m.table)

	var resp []*GroupRequests
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	return resp, err
}

func (m *defaultGroupRequestsModel) Insert(ctx context.Context, data *GroupRequests) (sql.Result, error) {
	groupRequestsIdKey := fmt.Sprintf("%s%v", cacheGroupRequestsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, groupRequestsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ReqId, data.GroupId, data.ReqMsg, data.ReqTime, data.JoinSource, data.InviterUserId, data.HandleUserId, data.HandleTime, data.HandleResult)
	}, groupRequestsIdKey)
	return ret, err
}

func (m *defaultGroupRequestsModel) Update(ctx context.Context, session sqlx.Session, data *GroupRequests) error {
	groupRequestsIdKey := fmt.Sprintf("%s%v", cacheGroupRequestsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupRequestsRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.ReqId, data.GroupId, data.ReqMsg, data.ReqTime, data.JoinSource, data.InviterUserId, data.HandleUserId, data.HandleTime, data.HandleResult, data.Id)
	}, groupRequestsIdKey)
	return err
}

func (m *defaultGroupRequestsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGroupRequestsIdPrefix, primary)
}

func (m *defaultGroupRequestsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRequestsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupRequestsModel) tableName() string {
	return m.table
}
