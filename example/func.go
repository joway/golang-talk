package main

import (
	"fmt"
)

func normal(msg string) (int, int) {
	fmt.Println(msg)
	return 0, 0
}

func multiArgs(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	var a, b = normal("Hi")
	fmt.Println(a, b)

	multiArgs("aaa", "bbb", "ccc")
}
