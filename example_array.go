//数组AND切片
package main

import (
	"fmt"
)
func main() {
	test_array_a()
	test_array_c()
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
	//var arrAge = [5]int{18, 20, 15, 22, 16}
	//var arrLazy = [...]int{5, 6, 7, 8, 22}
	//var arrLazy = []int{5, 6, 7, 8, 22}
	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	//var arrKeyValue = []string{3: "Chris", 4: "Ron"}
}

//把一个大数组传递给函数会消耗很多内存。2种方法可以避免这
//传递数组的指针(一般使用切片)2使用数组的切片
func test_array_c()  {
	//切片在未初始化之前默认为 nil，长度为 0。
	var arr1 = [5]int{1,2,3,4,5}
	var identifier []int//定义
	var identifier2 []int
	//初始化切片 var slice1 []type = arr1[start:end]
	identifier = arr1[0:2]//初始化
	identifier2 = arr1[:]
	for i:=0;i<len(identifier);i++ {
		fmt.Printf("slice at %d is %d",i,identifier[i])
	}
	for i:=0;i<len(identifier2);i++ {
		fmt.Printf("slice at %d is %d",i,identifier2[i])
	}
}
//切片slice是对数组的的一个连续片段的引用(该数组我们称之为相关的数组,通常是匿名的),所以切片是一个引用类型.
//切片的一些特性:1.切片是可索引的,可由len()函数获取长度,给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修
//改，最小为 0 最大为相关数组的长度：切片是一个 长度可变的数组。
//cap()切片的最大容量,0 <= len(s) <= cap(s)
//声明:var identifier []type 初始化:var slice1 []type = arr1[start:end] var slice1 []type = arr1[:] var slice1 []type = arr1[:]
//slice1 = slice1[:len(slice1)-1] s := []int{1,2,3} s := [3]int{1,2,3} s2 := s[:] s = s[:cap(s)]