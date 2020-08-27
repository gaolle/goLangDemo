package main

import (
	"fmt"
	"testing"
	"sync"
	"sync/atomic"
)

var total struct{
	sync.Mutex
	value int
	val uint64
}

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++{
		//加锁
		//total.Lock()
		//total.value += i
		//total.Unlock()
		atomic.AddUint64(&total.val, i)
	}
}

type TB struct {
	testing.TB
}
//重新实现 Fatal,隐式转换接口类型
func (p *TB) Fatal(args ...interface{})  {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("hello, playground")

	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}
