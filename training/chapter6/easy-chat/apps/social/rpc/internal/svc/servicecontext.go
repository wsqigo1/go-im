package svc

import (
	"github.com/wsqigo/easy-chat/apps/social/rpc/internal/config"
	"github.com/wsqigo/easy-chat/apps/social/socialmodels"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	socialmodels.FriendsModel        // 好友模型
	socialmodels.FriendRequestsModel // 好友申请模型
	socialmodels.GroupsModel         // 群模型
	socialmodels.GroupRequestsModel  // 群申请模型
	socialmodels.GroupMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		FriendsModel:        socialmodels.NewFriendsModel(sqlConn, c.Cache),
		FriendRequestsModel: socialmodels.NewFriendRequestsModel(sqlConn, c.Cache),
		GroupsModel:         socialmodels.NewGroupsModel(sqlConn, c.Cache),
		GroupRequestsModel:  socialmodels.NewGroupRequestsModel(sqlConn, c.Cache),
		GroupMembersModel:   socialmodels.NewGroupMembersModel(sqlConn, c.Cache),
	}
}
