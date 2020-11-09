package main

import (
	"fmt"
)

func main()  {
	// 字符串是所有8位字节字符串的集合，通常是但不一定表示UTF-8编码的文本。
	// 字符串可以为空，但不能为nil。字符串类型的值是不可变的。
	str := "hello"
	// str[0] = 'a' // Cannot assign to str[0]
	// Go语言字符串底层数据也是对应的字节数组，但是字符串的只读属性禁止了在程序中对底层字节数组的元素的修改。
	// 字符串赋值只是复制了数据地址和对应的长度，而不会导致底层数据的复制
	str1 := str
	// 字符串作为参数传递给fmt.Println函数时，字符串的内容并没有被复制
	// 传递的仅仅是字符串的地址和长度
	fmt.Println(str1)

	// 字符串的赋值操作也就是reflect.StringHeader结构体的复制过程，并不会涉及底层字节数组的复制
	// type StringHeader struct {
	// 	Data uintptr // 字符串指向的底层字节数组
	// 	Len  int     // 字符串的字节的长度
	// }

	// 符文是int32的别名，在所有方面都等同于int32。
	// 按照惯例，它用于区分字符值和整数值。
	// type rune = int32
}
