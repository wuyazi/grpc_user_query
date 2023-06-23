package logic

import (
	"context"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wuyazi/grpc_user_domain/user_domain"
	"github.com/wuyazi/grpc_user_query/internal/svc"
	"github.com/wuyazi/grpc_user_query/user_query"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUserJsonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUserJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUserJsonLogic {
	return &InsertUserJsonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	sqlGetUserJsonByUserId = `select user_id,json_data,create_time,update_time from user_json where user_id = ?`
	insertNewUserJsonSql   = `insert into user_json(user_id,json_data,create_time,update_time) values (?,?,?,?)`
)

func (l *InsertUserJsonLogic) InsertUserJson(in *user_domain.UserCreated) (*user_query.UserResp, error) {
	jsonString := protojson.Format(in.ProtoReflect().Interface())
	_, err := l.svcCtx.ReadDB.Exec(insertNewUserJsonSql, in.Abs.AggregateId, jsonString, in.CreateTime.AsTime(), in.CreateTime.AsTime())
	if err != nil {
		return nil, err
	}
	return &user_query.UserResp{}, nil
}
