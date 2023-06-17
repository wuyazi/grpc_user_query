package logic

import (
	"context"

	"github.com/wuyazi/grpc_user_domain/user_domain"
	"github.com/wuyazi/grpc_user_query/internal/svc"
	"github.com/wuyazi/grpc_user_query/user_query"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserLogic {
	return &InsertUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertUserLogic) InsertUser(in *user_domain.UserCreated) (*user_query.UserResp, error) {
	// todo: add your logic here and delete this line

	return &user_query.UserResp{}, nil
}
