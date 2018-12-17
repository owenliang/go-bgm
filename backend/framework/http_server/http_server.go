package http_server

import (
	"net/http"
	"net"
	"time"
	"github.com/owenliang/go-bgm/backend/framework/config"
	"strconv"
	"github.com/owenliang/go-bgm/backend/mvc"
)

type HttpServer struct {
	httpSever *http.Server
}

var (
	G_HttpServer *HttpServer
)

// HTTP服务
func InitHttpServer() (err error) {
	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
		staticDir http.Dir	// 静态文件根目录
		staticHandler http.Handler	// 静态文件的HTTP回调
		pattern string
		handleFunc http.HandlerFunc
	)

	// 配置路由
	mux = http.NewServeMux()
	for pattern, handleFunc = range mvc.Routes {
		mux.HandleFunc(pattern, handleFunc)
	}

	// css/js/img文件
	staticDir = http.Dir(config.G_Conf.StaticDir)
	staticHandler = http.FileServer(staticDir)
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	// 启动TCP监听
	if listener, err = net.Listen("tcp", ":" + strconv.Itoa(config.G_Conf.HttpPort)); err != nil {
		return
	}

	// HTTP服务
	httpServer = &http.Server{
		ReadTimeout: time.Duration(config.G_Conf.HttpReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.G_Conf.HttpWriteTimeout) * time.Millisecond,
		Handler: mux,
	}

	// 赋值单例
	G_HttpServer = &HttpServer{
		httpSever: httpServer,
	}

	// 启动
	go httpServer.Serve(listener)

	return
}