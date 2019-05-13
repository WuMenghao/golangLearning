package mian

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
·变量类型写在变量名之后
·编译器可以推测变量类型
·没有char,只有rune(4byte)
·原生支持复数类型


一、变量定义
1.var 关键字用于定义变量
2.var 关键词可以连续定义多个变量
3.var 关键词可以用 := 代替定义局部变量
4.var 关键词可以在包内方法外使用，:=不可以，在包内定义的变量只能在包内使用

二、变量类型
1.bool,string
2.(u)int ,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
3.byte,rune(字符，长度为32位，4字节)
4.float32,float64,complex64,complex128(复数)

三、类型转换
1.int()
2.类型转换是强制的

四、常量
1.const 关键字用于定义常量
2.const 数值可以不规定类型作为各种类型使用 const a,b = 3 ,4

五、枚举
1.普通枚举类型
	const (
		java = 0
		cpp = 1
		python = 2
		golang = 3
	)
2.自增值枚举类型
	const (
		java = iota //自增
		cpp
		python
		golang
	)
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Println("Hello word")
	fmt.Printf("%d %q \n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

var (
	aa = 3
	ss = "kkk"
	bb = true
)

//复数-欧拉公式
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	//保留3位
	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1)
}

//强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(fileName, c)
}

func enums() {
	//golang 没有枚举关键词，直接使用const定义一组常量
	const (
		//java = 0
		//cpp = 1
		//python = 2
		//golang = 3

		java = iota //自增
		cpp
		python
		golang
	)
	fmt.Println(java, cpp, python, golang)

	//iota 可用于计算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, pb, tb, pb)
}

func main() {
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()

	consts()
	enums()
}
