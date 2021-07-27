package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/types"
	"github.com/zhuxinlei/micro_user/user/model/common"
	"github.com/zhuxinlei/micro_user/user/model/input"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}

	password := string(hash)
	now := time.Now()
	data := input.User{
		Id:        0,
		UserName:  req.Username,
		Birthday:  req.Birthday,
		Gender:    req.Gender,
		Password:  password,
		CreatedAt: &now,
	}

	err = l.svcCtx.User.Insert(data)

	if err != nil {
		logx.Errorf("插入数据库错误", err.Error())
		return nil, errors.Wrapf(common.InsertUserError, common.ErrCodeMap[common.InsertUserErrorCode])
	}
	return &types.Reply{
		Code:    200,
		Message: "register success",
	}, nil
}
