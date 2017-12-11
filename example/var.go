package main

import "fmt"

func main() {
	// 先声明，后赋值
	var a string
	a = "a"

	// 声明并赋值，自带类型推导
	var b = "b"

	// 短声明变量，自带类型推导
	c := "c"

	fmt.Println(a, b, c)
}
