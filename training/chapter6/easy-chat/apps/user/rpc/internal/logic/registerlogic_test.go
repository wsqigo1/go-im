package logic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/wsqigo/easy-chat/apps/user/rpc/user"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	testCases := []struct {
		name     string
		req      *user.RegisterReq
		wantResp *user.RegisterResp
		wantErr  error
	}{
		{
			name: "1",
			req: &user.RegisterReq{
				Phone:    "19124155294",
				Nickname: "强盛老师",
				Avatar:   "png.jpg",
				Password: "123456",
				Sex:      1,
			},
			wantResp: &user.RegisterResp{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tc.req)
			assert.Equal(t, tc.wantErr, err)
			t.Log(got)
		})
	}
}
