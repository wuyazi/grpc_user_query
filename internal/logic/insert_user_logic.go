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

const (
	sqlGetUserByUserId      = `select id, nickname, age, gender from user where id = ?`
	insertNewUserProfileSql = `insert into user(id,nickname,create_time) values (?,?,?)`
)

func (l *InsertUserLogic) InsertUser(in *user_domain.UserCreated) (*user_query.UserResp, error) {
	ur := userRow{}
	_, err := l.svcCtx.ReadDB.Exec(insertNewUserProfileSql, in.Abs.AggregateId, in.Nickname, in.CreateTime.AsTime())
	if err != nil {
		return nil, err
	}
	userResp := renderRowUserResp(ur)
	return &userResp, nil
}

type userRow struct {
	Id       int64  `db:"id"`
	Nickname string `db:"nickname"`
	Age      int64  `db:"age"`
	Gender   int64  `db:"gender"`
}

func renderRowUserResp(ur userRow) user_query.UserResp {
	userResp := user_query.UserResp{}
	userResp.UserId = ur.Id
	userResp.Nickname = ur.Nickname
	userResp.Age = ur.Age
	switch ur.Gender {
	case 1:
		userResp.Gender = "MALE"
	case 2:
		userResp.Gender = "FEMALE"
	default:
		userResp.Gender = "UNKNOW"
	}
	return userResp
}
