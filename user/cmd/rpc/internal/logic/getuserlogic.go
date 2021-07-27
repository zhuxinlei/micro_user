package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_user/user/cmd/rpc/internal/svc"
	"github.com/zhuxinlei/micro_user/user/cmd/rpc/user"
	"github.com/zhuxinlei/micro_user/user/model/common"

	"github.com/zhuxinlei/micro_user/user/model/input"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("我是user rpc server 获取用户信息的业务逻辑")
	data := input.User{
		Id: int(in.Id),
	}
	data, err := l.svcCtx.User.GetUser(data)
	if err != nil {
		return nil, errors.Wrapf(common.GetUserError, common.ErrCodeMap[common.GetUserErrorCode])
	}

	return &user.UserInfoReply{
		Id:                   int64(data.Id),
		Name:                 data.UserName,
		Number:               "",
		Gender:               "",
	}, nil
}
