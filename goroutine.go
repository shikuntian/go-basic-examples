package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		go Add(i, i)
	}

}

func Add(x, y int) {
	fmt.Println(x)
}
