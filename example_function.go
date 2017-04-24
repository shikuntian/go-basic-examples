package main

import (
	"fmt"
)

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
	//test_fuc()
	//test_funa()
	//test_fun_f()
	//test_fun_m()
	//test_fun_o(2, 3)
	//test_fun_p(123456, test_fun_q)
	//test_fun_r();
	test_fun_s()
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
//关于参数传递,Go默认使用按值传递来传递参数,传递的是参数的副本,函数内部修改并不会改变其值
//如果加&,则传递的是一个指针地址,则修改值,参数本身的值也会改变,看example
func test_funa() {
	var a = 12;
	var xx = "the world";
	test_fun_d(a, &xx);
	fmt.Printf("\nthe input param value is%d\n", a)
	fmt.Print(xx)
}
func test_fun_d(c int, a *string) int {
	c = c + 1;
	fmt.Printf("the input param value is%d\n", c)
	*a = *a + "hello"
	fmt.Print(*a)
	return c + 1;
}
//跟PHP一样哈哈
//空白符号,空白符可以用来匹配一些不需要的值,然后丢掉,
//Q:如果丢掉,那不返回就好了,场景是什么?
func test_fun_f() {
	var int1 int
	var float float32
	int1, _, float = test_fun_i()
	fmt.Print(int1)
	fmt.Print(float)
}

func test_fun_i() (int, int, float32) {
	return 5, 6, 7.98
}

func test_fun_g(a int, b int) (min int, max int) {
	min = a + b
	max = b * a + b
	return
}

//改变外部的变量:传递指针
//传递变长参数
//概念:
//变参函数:如果一个函数最后一个参数采用...type的形式,那么这个函数就可以处理一个变长的参数,这个长度可以为0.
func test_fun_n(a int) {
	fmt.Printf("the first param is %d\n", a)
	test_fun_h(4, 2, 3, 4, 6)
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 7, 19, 21}
	test_fun_h(2, arr1...)
}
func test_fun_h(a int, len1 ...int) {
	fmt.Printf("the first param is %d\n", a)
	for i := range len1 {
		fmt.Printf("the second param is %d\n", i)
	}
}
//如果参数被存储在一个数组 arr 中，则可以通过 arr... 的形式来传递参数调用变参函数

//defer和追踪
//defer的作用就是让某个语句或函数在return前执行,
//defer可以传参数,多个defer的执行顺序是逆序执行
func test_fun_m() {
	fmt.Print("this is a begin")
	defer test_fun_n(1)
	defer fmt.Print("this is defer begin")
	fmt.Print("the is a end")
	//输出为this is a begin the is a end然后是执行test_fun_i
}

//defer使用场景:1.关闭文件流2.解锁一个加锁的资源3.打印最终报告4关闭数据库连接
//递归函数:在函数内部调用函数自己,就是递归
func test_fun_o(a int, b int) {
	var c = a + b
	fmt.Printf("the param's sum is:%d\n", c)
	if c < 12000 {
		test_fun_o(a, c)
	}
}
//6.7将函数作为参数
func test_fun_q(a int, b int) {
	fmt.Printf("the %d sqaquare is:%d", a, b)
}
func test_fun_p(a int, f func(int, int)) {
	fmt.Printf("the first param is:%d\n", a)
	f(a, a * a)
}
//闭包
func test_fun_r() int {
	fplus := func(x, y int) int {
		return x + y
	}
	fmt.Printf("the Closure Result is:%d and %v \n", fplus(2, 5),fplus)
	//the Closure Result is:7 and 0x48a600
	//fplus表示内存地址
	func() {
		var sum = 0.0
		for i := 1.0; i <= 1e6; i++ {
			sum += i
		}
		fmt.Printf("the sum is %d\n", sum)
	}()
	return fplus(12,12)
}
//应用闭包:将函数作为返回值
func test_fun_s()  {
	f:=func()func()int {
		return test_fun_r
	}
	fmt.Printf("the function result is:%d\n",f())
}

