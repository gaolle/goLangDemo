package main

import "fmt"

//删除9
func TrimSpace(arr []int) []int {
	b := arr[:0]
	for _, i := range arr{
		if i != 9{
			b = append(b, i)
		}
	}
	return b
}

func main()  {
	/*
	比数组多 Cap int(对应元素的个数的最大容量) 动态数组
	从0开始
	空切片 len cap all zero
	cap
	切片不会复制底层的数据
	*/
	var a[] int
	a = []int{2, 3, 5, 7, 11}
	c := a[1:3]
	c = append(c, []int{1, 2, 3, 4}...)
	c = append([]int{9, 9, 8}, c...)
	c = append(c, 0)
	copy(c[6:], c[5:])
	c[5] = 10
	c = c[5:]
	//原地操作删除
	c = append(c[:1], c[2:]...)

	//c = append(c[:1], append([]int{0, 0}, a[1:]...)...)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(c))
	fmt.Println(cap(c))
	for i, j := range c{
		fmt.Println(i, j)
	}

}
