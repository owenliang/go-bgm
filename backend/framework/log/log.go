package log

import (
	"k8s.io/klog"
	"flag"
)

// 初始化日志
func InitLog() {
	// 然后注册klog库需要的命令行参数
	klog.InitFlags(nil)

	// 强制覆盖klog的部分命令行参数
	flag.Set("stderrthreshold", "4") // 禁止klog输出任意级别日志到终端
}
