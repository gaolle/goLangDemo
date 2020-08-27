package main

import (
	"fmt"
	"sync"
	"time"
)

var do = make(chan bool) // 不带缓冲可运行

// var do = make(chan bool, 1) // 带缓冲不能运行，后台<- do 接收不到
var msg string

func aGorouine() {
	msg = "b"
	//<- do
	do <- true
}

/*
生产者消费者
*/
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i *factor
	}
}

func Consumer(in <-chan int)  {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	/*
		不带缓冲（同步原语），不取值将阻塞 接收发生在发生完成之前
	*/
	done := make(chan int)

	go func() {
		println("a")
		done <- 1
	}()
	<- done

	go aGorouine()
	//do <- true
	//do <- true
	<- do
	println(msg)

	/*
	多线程
	*/
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) //添加等待事件

		go func(i int) {
			println(i)
			wg.Done() //事件完成
		}(i)
	}
	wg.Wait() //等待完成

	/*
	生产者消费者
	*/
	ch := make(chan int, 8)

	go Producer(3, ch)
	go Producer(5, ch)
	//go Consumer(ch)

	time.Sleep(time.Second)
}
