package main

import (
	"fmt"
)

func chann(c, c1 chan int) {
	// 匹配时，执行匹配，当多个执行匹配时，随机选择执行其中一个,解决饥饿问题，避免存在某一个一直得不到执行
	select {
	case c <- 1:
		fmt.Println("chan c")
	case c1 <- 1:
		fmt.Println("chan c1")
	}
}

func main() {
	c := make(chan int)
	// 当存在default时，如果case不匹配，直接执行default
	select {
	case c <- 1:
		fmt.Println("chan")
	default:
		fmt.Println("default")
	}
	// default
}
