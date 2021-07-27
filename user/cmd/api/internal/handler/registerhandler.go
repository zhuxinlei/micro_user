package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/logic"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/response"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func registerHandler(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := types.RegisterReq{}
		/*if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}*/
		c := &gin.Context{
			Request:r,
		}
		err := c.ShouldBindJSON(&req)

		if err != nil{
			fmt.Println("参数错误",err.Error())
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {

			response.ServerError(w,err)

		} else {

			httpx.OkJson(w, resp)
		}
	}
}
