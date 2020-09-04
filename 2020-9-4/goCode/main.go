package main

import (
	"fmt"
)

// 通过接口实现多态
type AnimalInterface interface {
	Voice()
}

type Animal struct {
	Name string
	mean bool
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

/*
func (c *Cat) Voice() {
	if c.mean == true {
		fmt.Println("this is cat voice", c.Name)
	}
}

func (d *Dog) Voice() {
	if d.mean == true{
		fmt.Println("this is dog voice", d.Name)
	}
}
*/

func (c Cat) Voice() {
	if c.mean == true {
		fmt.Println("this is cat voice", c.Name)
	}
	c.AnimalVoice("this is cat voice by AnimalVoice")
}

func (d Dog) Voice() {
	if d.mean == true{
		fmt.Println("this is dog voice", d.Name)
	}
	d.AnimalVoice("this is dog voice by AniamlVoice")
}

func MakeSomeVoice(animalInterface AnimalInterface)  {
	animalInterface.Voice()
}

// 减少Voice重复性
func (animal Animal) AnimalVoice(voice string) {
	fmt.Println(voice)
}

func main() {
	dog := Dog{Animal{Name: "dog", mean: true}}
	cat := Cat{Animal{Name: "cat", mean: true}}
	// 值接收者和指针接收者(可以改变内部值)都可以调用，语言内部自己做了处理
	MakeSomeVoice(dog)
	MakeSomeVoice(cat)
}










