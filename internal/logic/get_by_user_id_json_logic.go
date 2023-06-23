package logic

import (
	"context"

	"github.com/wuyazi/grpc_user_query/internal/svc"
	"github.com/wuyazi/grpc_user_query/user_query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByUserIdJsonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByUserIdJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByUserIdJsonLogic {
	return &GetByUserIdJsonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByUserIdJsonLogic) GetByUserIdJson(in *user_query.GetByUserIdReq) (*user_query.UserResp, error) {
	ur := userRow{}
	err := l.svcCtx.ReadDB.QueryRow(&ur, sqlGetUserJsonByUserId, in.UserId)
	if err != nil {
		return nil, err
	}
	userResp := renderRowUserResp(ur)
	return &userResp, nil
}
