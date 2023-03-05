package main

import (
	"flag"
	"fmt"

	"monitor/monitor/internal/config"
	"monitor/monitor/internal/handler"
	"monitor/monitor/internal/logic"
	"monitor/monitor/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/monitor-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	config.RegisterProperty(ctx.Config.Monitor)

	logic.RegisterSchedule(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
