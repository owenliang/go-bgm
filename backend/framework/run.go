package framework

import (
	"github.com/owenliang/go-bgm/backend/framework/config"
	"flag"
	"runtime"
	"time"
	"github.com/owenliang/go-bgm/backend/framework/log"
	"github.com/owenliang/go-bgm/backend/framework/http_server"
	"github.com/owenliang/go-bgm/backend/framework/view"
)

var (
	confFile string
)

// 指定命令行参数
func initFlag() {
	flag.StringVar(&confFile, "config", "go-bgm.json", "go-bgm.json配置文件")
}

// 初始化线程
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 启动框架
func Run() (err error) {
	// go-bgm注册命令行参数
	initFlag()

	// 执行环境
	initEnv()

	// glog注册命令行参数
	log.InitLog()

	// 解析命令行参数
	flag.Parse()

	// 加载配置文件
	if err = config.InitConfig(confFile); err != nil {
		return
	}

	// 加载view
	if err = view.InitView(); err != nil {
		return
	}

	// 启动服务
	if err = http_server.InitHttpServer(); err != nil {
		return
	}

	for {
		time.Sleep(1 * time.Second)
	}

	return
}