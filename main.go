package main

import (
	"fmt"
	"go.uber.org/zap"
	"go_web_app/dao/mysql"
	"go_web_app/logger"
	"go_web_app/settings"
)

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings faild, err:%v\n", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger faild, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	// 3.初始化Mysql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4.初始化Redis连接
	// 5.注册路由
	// 6.启动服务
}
