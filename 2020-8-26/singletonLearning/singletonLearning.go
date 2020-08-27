package main

import (
	"sync"
)

//
type singleton struct {}
//
//var(
//	instance *singleton
//	initialized uint32
//	mu sync.Mutex
//)
//
//func Instance() *singleton {
//	if atomic.LoadUint32(&initialized) == 1 {
//		return instance
//	}
//
//	mu.Lock()
//	defer mu.Unlock()
//
//	if instance == nil {
//		defer atomic.StoreUint32(&initialized, 1)
//		instance = &singleton{}
//	}
//	return instance
//}

//type Once struct {
//	m sync.Mutex
//	done uint32
//}
//
//func (o *Once) Do(f func())  {
//	if atomic.LoadUint32(&o.done) == 1 { //加载数据
//		return
//	}
//
//	o.m.Lock()
//	defer o.m.Unlock()
//
//	if o.done == 0 {
//		defer atomic.StoreUint32(&o.done, 1) //保存数据
//		f()
//	}
//}
//
//var (
//	instance *singleton
//	once sync.Once
//)
//
//func Instance() *singleton  {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}

//顺序一致性
var str string
var done bool

func setup()  {
	str = "hello world"
	done = true
}

func hello()  {
	str = "hh"
	go f()
}

func f()  {
	println(str)
}

func main() {
	//初始化顺序 main.main之前都运行在一个主线程中
	//go setup()
	//for !done {}
	//println(str)
	//通道操作
	done := make(chan int)

	go func() {
		println("a")
		done <- 1
	}()
	<-done
	//互斥量操作
	var mu sync.Mutex
	i := 1
	mu.Lock()
	go func() {
		println(i)
		mu.Unlock()
	}()
	mu.Lock()

	go hello()

	//通道通信

}
