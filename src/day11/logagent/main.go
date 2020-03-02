package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"oldBoyGo/src/day11/logagent/config"
	"oldBoyGo/src/day11/logagent/kafka"
	"oldBoyGo/src/day11/logagent/tailf"
)

func main() {
	logPath := ".env"
	err := config.LoadConfig("ini", logPath)
	if err != nil {
		fmt.Println("load config failed, err:", err)
		panic("load config failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Println("init logger failed, err:", err)
		panic("init logger failed")
		return
	}
	logs.Debug("initialize success")
	logs.Debug("load config success, config%v", config.AppConfig)

	err = tailf.InitTailf(config.AppConfig.CollectsConfig, config.AppConfig.ChannelSize)
	if err != nil {
		logs.Error("init tailf failed, err:", err)
		return
	}

	err = kafka.InitKafka(config.AppConfig.KafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:", err)
		return
	}
	logs.Debug("initialize all success")

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("program exited")
}
