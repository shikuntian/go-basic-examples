package main

import "fmt"

func init() {

}
func main() {
	test_map_a()
}

//定义:map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元
//素是 value，所以这个结构也称为关联数组或字典
//声明方式:var map1 map[keytype]valuetype  var map1 map[string]int
//注意事项:1.未初始化的 map 的值是 nil,2.key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构
//体不能作为 key，但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方
//法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。3.value可以是任意类型,4.map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了
//多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中 直接读取要慢 100 倍；
func test_map_a() {
	var map1 map[int]string
	var map2 = map[string]int{"one":1,"two":12}
	map1 = map[int]string{1:"one", 2:"two"}
	fmt.Printf("the map1 key 1 is:%s\n",map1[1])
	fmt.Printf("the map2 key 1 is:%d\n",map2["one"])
}