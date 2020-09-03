### Go

25个关键字

import package （优先在GOROOT路径查找，然后在GOPATH查找）

— 空白标识符：抛弃不想继续使用的值

变量命名原则：不能以数字开头（保留关键字也不能用）

:= 左侧变量不能已经声明的，编译错误

if语句判断无括号

switch语句中的case自带break,当遇到fallthrough时，继续执行下一个case

指针指向一个变量的内存地址

new(type)：创建一个type类型的匿名变量，并初始化为type类型的零值，返回变量的地址，指针类型为*type

Map存储无序的键值对集合（底层Hash表的引用），底层数据的引用

```go
m := map[string]int {
	"a": 1,
	"b": 2,
	"c": 3,
}
m1 := m
delete(m, "a")
// 赋值时，底层数据是共享的
fmt.Println("m value is", m) // m value is map[b:2 c:3]
fmt.Println("m1 value is", m1) // m1 value is map[b:2 c:3]
```

```go
m := map[string]int {
	"a": 1,
	"b": 2,
	"c": 3,
}
m1 := map[string]int{}
// 循环赋值时，实现真正的拷贝（底层数据不共享）
for key, val := range m {
	m1[key] = val
}
delete(m, "a")
fmt.Println("m value is", m) // m value is map[b:2 c:3]
fmt.Println("m1 value is", m1) // m1 value is map[a:1 b:2 c:3]
```

切片、函数、包含切片的结构类型不能作为key(具有引用语义)，key唯一

结构体匿名成员

```go
// 只指定成员类型，不指定成员名，自动将成员类型作为成员名
type Week strcu {
    string
    int
    bool
}
```

函数允许返回多值

命名返回值

```go
func add(a, b int) (sum int) {
	sum = a + b
	return
}
```

匿名函数

``` go
var add = func(a, b int) (sum int) {
    sum = a + b
    return
}

func main() {
    fmt.Println(add(1, 2))
}
```

```go
// 匿名函数直接赋值
func main() {
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(sum)
}
```

可变参数函数

```go
// 将可变参数转换为新创建的切片
func change(str string, s ...string)  {
	if len(s) != 0 {
		s[0] = "d"
	}
	fmt.Printf("%T\n", s) // []string
	// 值改变，切片使用...解包时，将切片作为参数直接传入，不在新创建一个切片
	fmt.Println(s) // [d b c]
}

func main() {
	slice := []string{"a", "b", "c"}
	// 可变参数只接收string类型，而切片类型为[]string
	// change("", slice) // Cannot use 'slice' (type []string) as type string
	// 切片使用...解包，直接传入可变参数函数，不在创建新的切片
	change("", slice...)
}
```

方法（属于某一种类型，且有接收者），类型和方法处于同一包内 

```go
type Use struct {
	a, b int
}

// 将add绑定到Use类型 Use类型不能是接口类型或者接口指针列类型
func (u Use) add() int {
	return u.a + u.b
}

func main()  {
	u := Use{
		a: 1,
		b: 2,
	}
    // 通过Use类型调用add方法
	fmt.Println(u.add()) // 3
}
```

值接收者不会影响调用者的值，指针接收者会影响到调用者的值

```go
type Use struct {
	a, b int
}

// 值调用
func (u Use) modifedA(val int) {
	u.a = val
}

// 指针调用
func (u *Use) modifedB(val int) {
	u.b = val
}

func main()  {
	u := Use{
		a: 1,
		b: 2,
	}
	fmt.Println("值调用之前", u) // 值调用之前 {1 2}
	u.modifedA(4)
	fmt.Println("值调用之后", u) // 值调用之后 {1 2}

	fmt.Println("值指针调用之前", u) // 值指针调用之前 {1 2}
    // (&u).modifedB(4)
	u.modifedB(4)
	fmt.Println("指针调用之后", u) // 指针调用之后 {1 4}
}
```

嵌套使用

```go
type Student struct {
	name string
	age int
}

type School struct {
	Student // 匿名使用，两种方法调用
	index int
}

func (s *Student) modifedAge(age int) {
	s.age = age
}

func main() {
	school := School{
		index: 1,
		Student: Student{
			name: "a",
			age: 18,
		},
	}
	school.modifedAge(20)
	// school.Student.modifedAge(20)
	fmt.Println(school) // {{a 20} 1}
}
```

同名方法和同名函数的区别

```go
type Rect struct {
	width int
	height int
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 { // 同名方法可以定义在不同的类型上
	return float64(r.height * r.width)
}

// func (r Rect) Area() int { // 同名方法返回值不同也报错
// 	return r.height * r.width
// }

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// func qint() int { // 'qint' redeclared in this package
// 	return 1
// }
//
// func qint() float64 { // 同名函数报错
// 	return 1.0
// }

func main() {
	r := Rect{3, 4}
	c := Circle{3}
	fmt.Println("rect Area is", r.Area())
	fmt.Println("circle Area is", c.Area())
}
```

接口：一组方法的集合，不包含变量，不具体实现方法（接口定义方法，类型实现方法）

接口是一种类型，可以创建接口类型的变量

```go
type User struct {
	A, B int
}

// 类型和方法要在同一个包内 类型实现接口中的全部方法
func (u User) Add() int {
	return u.A + u.B
}

func (u User) Sub() int {
	return int(math.Abs(float64(u.B - u.A)))
}
```

```go
type Operator interface {
	Add() int
	Sub() int
}

func main() {
	var o Operator
    // 接口的动态类型变为User
	o = operator.User{A: 1, B: 2}
	fmt.Println("add", o.Add()) //add 3
	fmt.Println("sub", o.Sub()) //sub 1
}
```

一种类型可以实现多个接口

类型断言 

```go
// i为接口，Type是类型或接口 自动检测i的动态类型与Type是否一致
i.(Type) 
// val:Type所对应的值 ok:ture/false
val, ok := i.(Type)
```

接口不能实现别的接口也不能继承，通过嵌套接口创建新的接口

```go
// 嵌套接口
type Opts interface {
	Operator
}

type Operator interface {
	Add() int
	Sub() int
}

func main() {
	var opts Opts
	// 类型赋值给接口
	opts = operator.User{A: 1, B: 2}
	fmt.Println("add", opts.Add()) // add 3
}
```

Go语言通过携程实现并发，协程之间靠信道通信

协程：轻量级线程 在方法或函数调用之前加上关键字 go

信道(信道默认阻塞)

```
var c chan int
c := make(chan int)
c <- value
value := <-c
<-c
```

不能读写nil通道

```go
var c chan int
// goroutine 1 [chan send (nil chan)]:
// c <- 1
// <- c
```

单向通道

``` go
// 单向发送
var c chan<- int
// 单向接收
var c <-chan int
```

select没有输入值且只用于通道操作

```go
func send1(c chan string)  {
	c <- "send1"
}

func send2(c chan string)  {
	c <- "send2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go send1(c1)
	go send2(c2)

	// // select会阻塞
	// select {
	// case val := <- c1:
	// 	fmt.Println(val)
	// case val := <- c2:
	// 	fmt.Println(val)
	// }

	// select不会阻塞
	select {
	case val := <- c1:
		fmt.Println(val)
	case val := <- c2:
		fmt.Println(val)
	default:
		fmt.Println("no chan value")
	}
}
```

