package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

func main()  {
	/*
	函数包依次导入，不一次性直接包含
	具名函数（匿名函数的一种特例）
	匿名函数
	传值方式
	闭包：匿名引用外部变量
	*/
	var a = []interface{}{123, "abc"}
	Print(a...)
	Print(a)
	b := Inc()
	fmt.Println(b)

	//闭包延迟
	for i := 0; i < 3; i++ {
		//defer func() { println(i)}()
		var j = i
		defer func() {println(j)}()
		//defer func(i int) {( println(i))}(i)
	}
	/*
	方法关联到类型,编译时静态绑定的
	defer 延迟绑定
	*/

	var cp ColorPoint
	cp.Point.Y = 100
	cp.X = 10
	fmt.Println(cp.Point.X)
	fmt.Println(cp.Point.Y)
}

//解包
func Print(a ...interface{})  {
	fmt.Println(a...)
}
//闭包
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}