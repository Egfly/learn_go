package main

import "fmt"

/**
channel:
1、类似于unix中的管道
2、先进先出（队列类似）
3、线程安全，多个goroutine同时访问，不需要加锁
4、channel是有类型的， int类型的channel只能放int数据

声明：
var 变量名 chan 类型

赋值：
channel <- value

获取：
variable =<- channel
*/

type Student struct {
	name string
}

func main() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := Student{name: "student1"}
	stuChan <- &stu

	var student01 interface{}
	student01 = <-stuChan
	fmt.Println(student01)

	var student02 *Student
	student02, ok := student01.(*Student)
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(student02)
}
