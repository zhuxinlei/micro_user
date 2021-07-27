package logic

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/types"
	"github.com/zhuxinlei/micro_user/user/model/common"

	"github.com/zhuxinlei/micro_user/user/model/input"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	user := input.User{
		UserName: req.Username,
	}
	data, err := l.svcCtx.User.GetUser(user)
	if err != nil {
		return nil, errors.Wrapf(common.GetUserError, common.ErrCodeMap[common.GetUserErrorCode])
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password)) //验证（对比）
	if err != nil {
		fmt.Println("pwd wrong")
		return nil, errors.Wrapf(common.GetUserError, common.ErrCodeMap[common.GetUserErrorCode])
	} else {
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(data.Id))
		if err != nil {
			return nil, err
		}
		loginReply := types.LoginReply{
			Id:           0,
			Name:         data.UserName,
			Gender:       "",
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}
		return &types.Reply{
			Code:    200,
			Message: "success",
			Data:    loginReply,
		}, nil

	}

}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
