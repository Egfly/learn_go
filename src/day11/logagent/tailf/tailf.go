package tailf

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"oldBoyGo/src/day11/logagent/config"
	"time"
)

var (
	tailObjList *TailObjMgr
)

type LogMsg struct {
	Msg   string
	Topic string
}

type TailObj struct {
	tail *tail.Tail
	conf config.CollectConfig
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *LogMsg
}

/**
从channel获取日志
*/
func GetLine() (msg *LogMsg) {
	msg = <-tailObjList.msgChan
	return
}

func InitTailf(config []config.CollectConfig, channelSize int) (err error) {
	if len(config) == 0 {
		err = errors.New("empty collect config")
		return
	}
	tailObjList = &TailObjMgr{
		msgChan: make(chan *LogMsg, channelSize),
	}
	for _, v := range config {
		obj := &TailObj{
			conf: v,
		}
		tails, tailErr := tail.TailFile(v.LogPath, tail.Config{
			//Location:  &tail.SeekInfo{Offset: 0, Whence: 1},
			ReOpen:    true,
			MustExist: false,
			Poll:      true,
			Follow:    true,
		})
		if tailErr != nil {
			err = tailErr
			return
		}
		obj.tail = tails
		tailObjList.tailObjs = append(tailObjList.tailObjs, obj)
		go readFromTail(obj)
	}
	return
}

func readFromTail(tails *TailObj) {
	for true {
		line, ok := <-tails.tail.Lines
		if !ok {
			logs.Warn("tail file reopen, filename:%s", tails.tail.Filename)
			time.Sleep(time.Second)
			continue
		}
		logMsg := &LogMsg{
			Msg:   line.Text,
			Topic: tails.conf.Topic,
		}
		tailObjList.msgChan <- logMsg
	}
}
