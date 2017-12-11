//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// 同步通信
//var pool = make(chan int, 1)
//
//// 同步状态
//var done = make(chan bool)
//
//func producer(num int) {
//	name := fmt.Sprintf("Producer-%d", num)
//	for {
//		msg := rand.Int()
//		pool <- msg
//		fmt.Printf("%s : sent %d\n", name, msg)
//	}
//}
//
//func consumer(num int) {
//	name := fmt.Sprintf("Consumer-%d", num)
//	for {
//		select {
//		case msg, ok := <-pool:
//			if !ok {
//				done <- true
//				return
//			}
//			fmt.Printf("%s : get %d\n", name, msg)
//		default:
//			time.Sleep(10 * time.Millisecond)
//		}
//	}
//}
//
//func main() {
//	producerCount := 2
//	consumerCount := 4
//
//	for i := 0; i < producerCount; i++ {
//		go producer(0)
//	}
//	for i := 0; i < consumerCount; i++ {
//		go consumer(i)
//	}
//
//	for {
//		<-done
//	}
//}
