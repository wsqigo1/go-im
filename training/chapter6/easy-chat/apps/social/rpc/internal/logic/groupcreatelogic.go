package logic

import (
	"context"
	"github.com/wsqigo/easy-chat/apps/social/rpc/internal/svc"
	"github.com/wsqigo/easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群要求
func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	// todo: add your logic here and delete this line
	//groups := &socialmodels.Groups{
	//	Id:         wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
	//	Name:       in.Name,
	//	Icon:       in.Icon,
	//	CreatorUid: in.CreatorUid,
	//	IsVerify:   false,
	//}
	//
	//err := l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
	//
	//})
	return &social.GroupCreateResp{}, nil
}
