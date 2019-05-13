package mian

import "fmt"

/**
一、指针
1.指针定义：
	var a int = 2
	var pa *int = &a
2.特点：
	不能用于运算

二、参数传递
1.golang只有值传递一种方式（函数式语言）

*/

//使用指针做换值操作
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
