// Code generated by goctl. DO NOT EDIT.
package types

import "time"

type Reply struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   uint8  `json:"gender"`
	Birthday time.Time `json:"birthday"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
