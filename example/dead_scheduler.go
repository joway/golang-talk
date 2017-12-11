package main

import (
	"time"
	"fmt"
)

func dead() {
	i := 0
	for i < 0 {
		i++
	}
}

func main() {
	go dead()

	for {
		fmt.Println("1")
		time.Sleep(time.Second)
	}
}
