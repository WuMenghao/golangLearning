package mian

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
一、函数
1.函数可以返回多个值
2.函数返回多个值时可以起名
3.仅用于非常简单的函数
4.对于调用者而言没有区别（对业务而已有区分意义）
5.函数的参数、函数体内、函数的返回值 都可以是函数

二、特点
1.无默认参数 2.无重载 3.有可变参数

三、总结
1.返回值类型写在最后面
2.返回值可以有多个
3.函数作为参数
4.没有默认参数，可选参数
*/

//四则运算
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//除余计算
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	//return
	return a / b, a % b
}

//函数参数为函数
func apply(op func(int, int) int, a, b int) int {
	//获取函数名 1.反射获取指针 2.获取函数信息中的名称
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d ,%d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(13, 3))

	//多返回值
	q, r := div(13, 3)
	fmt.Println(q, r)

	//多个返回值只接受其中一个
	i, _ := div(13, 3)
	fmt.Println(i)

	//异常及处理
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	//函数参数为函数
	fmt.Println(apply(pow, 3, 4))

	//函数参数传入匿名函数
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	//可变参数
	fmt.Println(sum(1, 2, 3, 4, 5))
}
