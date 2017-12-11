package main

import (
	"fmt"
	"math/rand"
)

// 同步通信
var pool = make(chan int, 1)

// 同步状态
var done = make(chan bool)

func producer(num int) {
	name := fmt.Sprintf("Producer-%d", num)
	for {
		msg := rand.Int()
		pool <- msg
		fmt.Printf("%s : sent %d\n", name, msg)
	}
}

func consumer(num int) {
	name := fmt.Sprintf("Consumer-%d", num)
	for msg := range pool {
		fmt.Printf("%s : get %d\n", name, msg)
	}
	done <- true
}

func main() {
	producerCount := 2
	consumerCount := 4

	for i := 0; i < producerCount; i++ {
		go producer(i)
	}
	for i := 0; i < consumerCount; i++ {
		go consumer(i)
	}

	for {
		<-done
	}
}
