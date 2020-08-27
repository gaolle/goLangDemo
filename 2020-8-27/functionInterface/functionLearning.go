package main

import (
	"fmt"
	"image/color"
	"testing"
)
/*
方法关联到类型，编译阶段完成方法静态绑定
*/
type File struct {
	fd int
}
/*
通过将类型名前提，让函数与类型绑定，不在于类型声明的对象只绑定
*/
func closeFile(f *File)  {
}

func (f *File) CloseFile() {
}

func (f *File) ReadFile()  {
}
/*
结构体内置匿名成员实现继承，不能实现虚函数多态特性
*/
type Point struct {
	X, Y float64
}

type ColoedPoint struct {
	Point //匿名成员 内嵌Point成员X，Y
	Color color.RGBA
}
/*
虚函数多态接口实现：延迟绑定
*/
type TB struct {
	testing.TB //匿名成员
}
//类型重新实现接口
func (p *TB) Fatal(args ...interface{})  {
	fmt.Println("TB")
}

func main() {
	var cp ColoedPoint //通过匿名继承Point成员、方法，接收者依然是匿名成员，不是cp当前变量
	cp.X = 10
	fmt.Println(cp.Point.X)

	var tb testing.TB = new(TB) //隐式转换为匿名成员接口类型
	tb.Fatal("tb")		//调用新的接口 打印TB
}
