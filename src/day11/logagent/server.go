package main

import (
	"github.com/astaxie/beego/logs"
	"oldBoyGo/src/day11/logagent/kafka"
	"oldBoyGo/src/day11/logagent/tailf"
	"time"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}

	return
}

func sendToKafka(msg *tailf.LogMsg) (err error) {
	err = kafka.Send(msg.Msg, msg.Topic)
	return
}
