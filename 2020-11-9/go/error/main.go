package main

import (
	"fmt"
	"sync"
)

// WaitGroup类型有三个方法（特殊的函数，将在以后的文章中详解）：Add、Done和Wait。 此类型将在后面的某篇文章中详细解释，目前我们可以简单地认为：
// Add方法用来注册新的需要完成的任务数。
// Done方法用来通知某个任务已经完成了。
// 一个Wait方法调用将阻塞（等待）到所有任务都已经完成之后才继续执行其后的语句。

var wg sync.WaitGroup

func Sync() {
	fmt.Println(1)
	wg.Done()
}

func main() {
	// 我们必须在defer函数中直接调用recover recover处于的defer 和panic处于同一goroutine
	// recover函数捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层defer函数）
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	//
	// panic(123)

	// Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU：
	
	wg.Add(2)
	go Sync()
	go Sync()
	wg.Wait()
}
