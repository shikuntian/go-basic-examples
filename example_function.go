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
	//test_fuc()
	//test_funa()
	test_fun_f()
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
func test_funa()  {
	var a = 12;
	var xx = "the world";
	test_fun_d(a,&xx);
	fmt.Printf("\nthe input param value is%d\n",a)
	fmt.Print(xx)
}
func test_fun_d(c int,a *string) int  {
	c = c+1;
	fmt.Printf("the input param value is%d\n",c)
	*a = *a+"hello"
	fmt.Print(*a)
	return c+1;
}
//跟PHP一样哈哈
//空白符号,空白符可以用来匹配一些不需要的值,然后丢掉,
//Q:如果丢掉,那不返回就好了,场景是什么?
func test_fun_f()  {
	var int1 int
	var float float32
	int1,_,float = test_fun_i()
	fmt.Print(int1)
	fmt.Print(float)
}

func test_fun_i()(int,int,float32)  {
	return 5,6,7.98
}

func test_fun_g(a int,b int)(min int,max int)  {
	min=a+b
	max=b*a+b
	return
}

//改变外部的变量:传递指针
