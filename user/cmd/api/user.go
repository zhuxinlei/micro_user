package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/config"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/handler"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	cc := logx.LogConf{
		ServiceName:         "user.api",
		Mode:                "file",
		TimeFormat:          "",
		Path:                "logs",
		Level:               "",
		Compress:            false,
		KeepDays:            0,
		StackCooldownMillis: 0,
	}
	logx.MustSetup(cc)
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
