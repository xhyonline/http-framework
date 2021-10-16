package main

import (
	"net/http"

	"github.com/xhyonline/http-framework/configs"
	"github.com/xhyonline/http-framework/internal"
	"github.com/xhyonline/http-framework/middleware"
	"github.com/xhyonline/http-framework/router"

	"github.com/xhyonline/xutil/sig"

	// nolint
	. "github.com/xhyonline/http-framework/component" // 忽略包名

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	// 初始化配置
	configs.Init(configs.WithRedis(), configs.WithMySQL())
	// 初始化 mysql 、redis 等服务组件
	ComponentInit(RegisterRedis(), RegisterMySQL(), RegisterLogger())
	// 中间件
	g.Use(middleware.Cors())
	// 初始化路由
	router.InitRouter(g)
	// 启动 HTTP 服务
	httpServer := &internal.HTTPServer{Server: &http.Server{Addr: ":8080", Handler: g}}
	go httpServer.Run()
	// 注册优雅退出
	ctx := sig.Get().RegisterClose(httpServer)
	<-ctx.Done()
}
