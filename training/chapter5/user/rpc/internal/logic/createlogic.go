package logic

import (
	"context"
	"github.com/wsqigo/training/docker/user/models"

	"github.com/wsqigo/training/docker/user/rpc/internal/svc"
	"github.com/wsqigo/training/docker/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Name:  in.Name,
		Phone: in.Phone,
	})
	return &user.CreateResp{}, err
}
