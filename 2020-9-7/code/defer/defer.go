package main

import (
	"fmt"
	time2 "time"
)

func main() {
	// defer 插入顺序从后往前，执行时从前往后 （栈的顺序）
	// // defer调用时刻：在函数或者方法返回之前调用
	// {
	// 	defer fmt.Println("defer run")
	// 	fmt.Println("block run")
	// }
	// fmt.Println("main run")
	// // block run
	// // main run
	// // defer run

	time := time2.Now()
	// 函数都是参数传值，函数拷贝传进去的值
	// defer fmt.Println(time2.Since(time)) // 0
	// 匿名函数对外部值是引用，拷贝的是函数指针
	defer func() {fmt.Println(time2.Since(time))}() //1.0009793s
	time2.Sleep(time2.Second)
}
