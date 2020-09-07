package main

import (
 "fmt"
)

func main() {
	// // panic中断程序，但会执行defer
	// defer func() {
	// 	// recover只能在defer中使用
	// 	recover()
	// 	fmt.Println("defer")
	// }()
	// // 当发生panic之后,会直接执行defer,后面代码不会继续执行
	// panic(1)
	// fmt.Println("main")
	// defer func() {
	// 	recover()
	// 	fmt.Println("defer 2")
	// }()
	// // defer

// 	// panic只会触发当前goroutine的延迟函数调用(defer)
// 	defer fmt.Println("main defer")
// 	go func() {
// 		defer fmt.Println("func defer")
// 		panic("func defer")
// 	}()
// 	time2.Sleep(time2.Second)
// 	// func defer
// 	// panic: func defer

    // panic嵌套执行顺序
	defer fmt.Println("main defer")
	defer func() {
	   defer func() {
	       fmt.Println("func func defer")
       }()
	   fmt.Println("func defer")
    }()
	panic("main panic")
	// func defer
	// func func defer
	// main defer
 }
