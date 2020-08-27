package main

import (
	"fmt"
	"time"
)

//带缓存不能打印  带缓冲 K个接收完成操作发生在k+c(channel的缓存大小)个发生操作之前
//var done = make(chan bool, 1)
//无缓冲 接收发生在发生完成之前
//var done = make(chan bool)
//var msg string
//
//func aGoroutine()  {
//	msg = "hello"
//	<-done
//}

var limit = make(chan int, 3)
var done = make(chan int, 10)
var work = []func(){
	func() {
		println("1"); time.Sleep(1 * time.Second)
	},
	func() {
		println("2"); time.Sleep(1 * time.Second)
	},
	func() {
		println("3"); time.Sleep(1 * time.Second)
	},
	func() {
		println("4"); time.Sleep(1 * time.Second)
	},
	func() {
		println("5"); time.Sleep(1 * time.Second)
	},
}

func main() {
	//go aGoroutine()
	//done <- true
	//println(msg)
	//for _, w := range work{
	//	go func(w func()) {
	//		limit <- 1
	//		w()
	//		<-limit
	//	}(w)
	//}
	////select {
	////}
	//time.Sleep(10 * time.Second)
	for i := 0; i < 10; i++ {
		//fmt.Println(", ", &i, i)
		go func(i int) {
			fmt.Println(&i)
		}(i)

	}
	// select{}
	time.Sleep(100 * time.Second)
}
