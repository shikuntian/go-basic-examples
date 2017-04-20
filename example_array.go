//数组AND切片
package main

import "fmt"

func main() {
	test_array_a()
}

//数组有特定的用处，但是却有一些呆板，所以在 Go 语
//言的代码里并不是特别常见,切片确实随处可见的数组长度最大为
//2Gb

func test_array_a()  {
	var int1 [10]int //长度为5的int类型的数组
	for i:=0;i<len(int1);i++ {
		int1[i] = i*i
	}
	for i:=0;i<len(int1);i++ {
		fmt.Printf("index is %d and value is %d\n",i,int1[i])
	}
	for i:= range int1{
		fmt.Printf("index is %d and value is %d\n",i,int1[i])
	}
}

func test_array_b()  {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrLazy = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrKeyValue = []string{3: "Chris", 4: "Ron"}
}

//把一个大数组传递给函数会消耗很多内存。2种方法可以避免这
//传递数组的指针(一般使用切片)2使用数组的切片
