package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
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

func main() {
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
}
