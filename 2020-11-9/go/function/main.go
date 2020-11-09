package main

import (
	"fmt"
)

// // 具名函数
// func Add(a, b int) int {
// 	return a+b
// }
//
// // 匿名函数
// var Add = func(a, b int) int {
// 	return a+b
// }

func main() {
	// 唯一一种引用值 闭包
	for i := 0; i < 2; i++ {
		// 延迟执行
		defer func() {
			// i 为引用传值
			fmt.Println(i)
		}()
	}

	for i:= 0; i < 2; i++ {
		defer func(int2 int) {
			fmt.Println(int2)
		}(i)
	}

	// 不用关心Go语言中函数栈和堆的问题，编译器和运行时会帮我们搞定；
	// 同样不要假设变量在内存中的位置是固定不变的，指针随时可能会变化，特别是在你不期望它变化的时候

}
