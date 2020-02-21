package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

/**
beego config package example
*/
func main() {
	// init conf
	conf, err := config.NewConfig("ini", "test.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	ipAddr := conf.String("server::ip_addr")
	if len(ipAddr) == 0 {
		fmt.Println("invalid ip addr")
	}
	fmt.Println("ip address: ", ipAddr)

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:ip_port failed, err:", err)
	}
	fmt.Println("port: ", port)

	logLevel := conf.String("log::log_level")
	fmt.Println("log level: ", logLevel)
}
