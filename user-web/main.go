package main

import (
	"fmt"

	"go.uber.org/zap"

	"yuan-api/user-web/initialize"
)

func main() {

	//初始化logger
	initialize.InitLogger()

	//初始化routers
	Router := initialize.Routers()

	port := 8021

	zap.S().Infof("启动服务器, 端口: %d", port)

	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panicf("启动失败, err: %s", err.Error())
	}

}
