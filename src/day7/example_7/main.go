package main

import (
	"flag"
	"fmt"
)

/**
命令行参数解析: flag package

flag.BoolVar(&test, "b", false, "print on newline")
flag.StringVar(&str, "s", "", "print on newline")
flag.IntVar(&count, "c", 1001, "print on newline")

*/
func main() {
	var configPath string
	var logLevel int
	// param1: 用来接收命令行参数的地址   param2: 参数key   param3: 参数默认值   param4: 参数缺失时的提示语
	flag.StringVar(&configPath, "path", "", "please input config path")
	flag.IntVar(&logLevel, "level", 10, "please input log level")

	// Parse从参数os.Args[1:]中解析命令行标签。 这个方法调用时间点必须在FlagSet的所有标签都定义之后，程序访问这些标签之前。
	flag.Parse()

	// ./example_7 -path /opt/www/index.php -level 20
	// path: /opt/www/index.php
	// logLevel: 20
	fmt.Println("path:", configPath)
	fmt.Println("logLevel:", logLevel)
}
