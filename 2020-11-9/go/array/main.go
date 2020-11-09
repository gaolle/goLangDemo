package main

import (
	"fmt"
)

func main() {
	// 一个数组变量即表示整个数组，它并不是隐式的指向第一个元素的指针（比如C语言的数组），而是一个完整的值。当一个数组变量被赋值或者被传递的时候，实际上会复制整个数组。
	// 如果数组较大的话，数组的赋值也会有较大的开销。
	// 为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组
	var arr = [...]int{1, 2, 3, 4}
	arrPtr := &arr // 指向数组指针

	// rang遍历不会导致越界
	for _, v := range arrPtr{
		fmt.Println(v)
	}

	// 长度为0的数组在内存不占空间
	// var a [0]int


}