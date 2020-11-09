package main

import (
	"fmt"
	"image/color"
)

// 每种类型对应的方法必须和类型的定义在同一个包中
// 对于给定的类型，每个方法的名字必须是唯一的，同时方法和函数一样也不支持重载

type Student struct {
	Name string
	Age int
}

func (stu *Student) setName() {
	stu.Name = "A"
}

func (stu *Student) setAge() {
	stu.Age = 18
}

// 通过在结构体内置匿名的成员来实现继承
type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}


func main() {
	// 方法关联到类型 在编译阶段完成方法的静态绑定
	stu := Student{}
	fmt.Println(stu)
	stu.setName()
	stu.setAge()

	fmt.Println(stu.Age)
	fmt.Println(stu.Name)
}
