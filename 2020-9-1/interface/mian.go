package main

import (
	"fmt"
)

// 接口是一组方法的集合
// 当一个类型提供了接口中方法的定义时,类型实现接口
// 接口也是一种类型,可以创建接口类型的变量
// 访问接口类型,返回的是底层动态值的类型
//
// type Iname interface {
// 	Mname()
// }
//
// type St struct {
// }
//
// func (St) Mname() {
// }
//
// type s struct {
//
// }
//
// func main() {
// 	var st *St
// 	if st == nil {
// 		fmt.Println("this is nil value")
// 	} else {
// 		fmt.Println("thi is no nil value")
// 	}
//
// 	// st 的动态值为nil,但动态类型为*St
// 	var i Iname = st
// 	fmt.Println(i)
// 	if i == nil {
// 		fmt.Println("this is nil value")
// 	} else {
// 		fmt.Println("thi is no nil value")
// 	}
// }

type Shape interface {
	Area() float32
}

type Rect struct {
	width float32
	height float32
}

func (r Rect) Area() float32 {
	return r.width * r.height
}

func main() {
	var s Shape
	s = Rect{width: 3, height: 5}
	fmt.Println("area", s.Area())
}


