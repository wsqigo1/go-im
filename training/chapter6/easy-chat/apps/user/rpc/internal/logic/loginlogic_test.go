package logic

import (
	"github.com/wsqigo/easy-chat/apps/user/rpc/internal/config"
	"github.com/wsqigo/easy-chat/apps/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config

	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
