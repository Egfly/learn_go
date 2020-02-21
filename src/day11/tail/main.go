package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

/**
tail package example
*/
func main() {
	filename := "test.txt"
	tails, err := tail.TailFile(filename, tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen. filename: %s \n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", msg.Text)
	}
}
