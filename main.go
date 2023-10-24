package main

import (
	"plan/config"
	"plan/route"
)

// 初始化配置，里面可以初始化mysql或redis等等
func init() {
	if err := config.InitMysql(); err != nil {
		panic(err)
	}
}

func main() {

	//注册路由并启动gin
	if err := route.InitRoute().Run(":8081"); err != nil {
		panic(err)
	}

}
