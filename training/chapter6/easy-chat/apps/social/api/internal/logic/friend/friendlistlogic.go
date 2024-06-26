package friend

import (
	"context"
	"github.com/wsqigo/easy-chat/apps/social/rpc/socialclient"
	"github.com/wsqigo/easy-chat/apps/user/rpc/userclient"
	"github.com/wsqigo/easy-chat/pkg/ctxdata"

	"github.com/wsqigo/easy-chat/apps/social/api/internal/svc"
	"github.com/wsqigo/easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxdata.GetUId(l.ctx)

	friends, err := l.svcCtx.Social.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 根据好友id获取好友信息
	uids := make([]string, 0, len(friends.List))
	for _, f := range friends.List {
		uids = append(uids, f.FriendUid)
	}

	// 根据uids查询用户信息
	users, err := l.svcCtx.User.FindUser(l.ctx, &userclient.FindUserReq{
		Ids: uids,
	})
	if err != nil {
		return &types.FriendListResp{}, nil
	}
	userRecords := make(map[string]*userclient.UserEntity, len(users.User))
	for i, user := range users.User {
		userRecords[user.Id] = users.User[i]
	}

	respList := make([]*types.Friends, 0, len(friends.List))
	for _, f := range friends.List {
		friend := &types.Friends{
			Id:        f.Id,
			FriendUid: f.FriendUid,
		}

		if u, ok := userRecords[f.FriendUid]; ok {
			friend.Nickname = u.Nickname
			friend.Avatar = u.Avatar
		}
		respList = append(respList, friend)
	}

	return &types.FriendListResp{
		List: respList,
	}, nil
}
