package main

import "fmt"

//定义一个初始化函数,如果有此函数
//会先执行它,然后执行main函数
//如果没有main函数,报错信息如下:
//runtime.main_main·f: relocation target main.main not defined
//runtime.main_main·f: undefined: "main.main"
//意思就是需要定义一个main函数
//貌似如果一个函数的参数没有指定类型,默认是int类型
func init() {

}

func main() {
	test_fuc()
}

func test_fuc() {
	fmt.Print("the test function is begin")
	test_fun_b(test_fun_a("11", 12))
}

func test_fun_a(a string, c int) (int, int, int) {
	fmt.Print(a)
	fmt.Print(c)
	return 1, 2, 3
}

func test_fun_b(a, b, c int) {
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
}

//go的函数可以返回多个值,比如上面的test_fun_a
//任何一个有返回值的函数都必须有return 或panic
//嘿嘿,有返回值当然是用return了,panic好像是捕获错误的
//函数必须满足函数的返回类型和参数数量 比如test_fun_a必须返回三个参数且都为整数
//关于参数传递,Go默认使用按值传递来传递参数,传递的是参数的副本,