package logic

import (
	"context"

	"github.com/wuyazi/grpc_user_query/internal/svc"
	"github.com/wuyazi/grpc_user_query/user_query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByUserIdLogic {
	return &GetByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc insertUser(user_domain.UserCreated) returns(userResp);
func (l *GetByUserIdLogic) GetByUserId(in *user_query.GetByUserIdReq) (*user_query.UserResp, error) {
	// todo: add your logic here and delete this line

	return &user_query.UserResp{}, nil
}
