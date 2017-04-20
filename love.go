package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	Unknown = 0
	Female = 1
	Male = 2
)

var sexual int
var smart bool = true;

var a8 int8
var a16 int16
var f32 float32
var f321 float32
var weiwei string
func go_app(item int) {
	fmt.Printf("I Love You So Much" + "\n")
	switch item {
	case Unknown:
		fmt.Printf("I don't know what i like" + "\n")
	case Female:
		fmt.Printf("I Like Female" + "\n")
	case Male:
		fmt.Printf("I like Male" + "\n")
	}
	//if else
	if smart {
		fmt.Printf("You are so smart" + "\n")
	} else{
		fmt.Print("You are a fool \n")
	}
	if !smart {
		fmt.Print("You are a fool \n")
	}
	fmt.Print("the int 8 number is \n", a8)
	//fmt.Print(a8+a16)
	fmt.Print(f32 + f321)
	fmt.Print("the weiwei length is:", len(weiwei))
	fmt.Print("\n the name is :", weiwei[5])
	s := "hel" + "lo"
	fmt.Print(s)
	fmt.Printf("%t \n",strings.HasPrefix(s,"Hel"))
	fmt.Printf("%t \n",strings.Contains(s,"el"))
	//时间
	te:=time.Now()
	fmt.Println(te)
	fmt.Printf("%02d.%02d.%4d\n",te.Day(),te.Month(),te.Year())
	te2:=time.Now().UTC()
	fmt.Println(te2)
	//指针
	var i1=5
	fmt.Printf("An integer:%d,Its location in memmory is:%p\n",i1,&i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	*intP = 6
	fmt.Printf("An integer:%d,Its location in memmory is:%p\n",i1,&i1)
	var p *int
	p = &i1
	fmt.Print(p)
}

func main() {
	go_app(sexual)
}

func init() {
	sexual = 0
	a8 = 12
	a16 = 32767
	f32 = 43.2222
	weiwei = "I Love You Forever"
}