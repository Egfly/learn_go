package main

import (
	test "day2/example_3/test_init"
	"fmt"
)

/**
1、包的别名写在包的路径前面
2、如果只想导入包，但是不使用包中的变量或者函数，需要将 _ 设置为包的别名
3、每个包必定包含一个init函数
4、init函数在全局变量赋值之后运行
5、在程序执行时，程序会先自下而上，层层导入所有包，导入时会触发init函数。当所有包导入完之后，才会执行main函数中的代码
6、自下而上导入：main =>(发现) test_init =>(发现) test_init2 =>(初始化) test_init2 =>(初始化) test_init =>(初始化) main
7、导入==初始化
*/
func main() {
	fmt.Printf("%s 的年龄是%d", test.Name, test.Age)
	fmt.Println()
	test.Test()
}
