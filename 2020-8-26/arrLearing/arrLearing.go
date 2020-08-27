package main

import "fmt"

func main()  {
	/*

	*/
	var a = [...]int{1, 2, 3}
	var b = &a	//整体操作
	fmt.Println(a[0], b[2])
	for _, i := range b { //类似键值 _ ： 不读取  不会数组越界
		fmt.Println(i)
	}
	var times [5][0]int
	for range times{ //二维不起作用 长度为0的数组在内存中不占空间
		fmt.Println("h")
	}
}