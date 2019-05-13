package main

import "fmt"

/**
一、数组的定义
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	grid := [4][5]int{}

二、range的使用
	1.使用range的原因：意义明确，美观
	2.range可以同时获得index 和 value

三、golang中数组是值类型
	1.[3]int 与 [5]int 是不同的
	2.调用func f(arr [10]int) 会拷贝数组
	3.一般不直接使用数组

*/

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayChange(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayWithPointer(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	//定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	grid := [4][5]int{}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历 range的使用
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		//下标和值
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		//可以通过_省略变量
		fmt.Println(v)
	}

	//最大值索引和值
	maxi := -1
	maxv := -1
	for i, v := range arr3 {
		if v > maxv {
			maxi, maxv = i, v
		}
	}
	fmt.Println(maxi, maxv)

	printArray(arr1)
	//printArray(arr2) //编译错误arr2是[3]int 无法传入 [5]int
	printArray(arr3)

	//证明是值传递
	fmt.Println("printArrayChange(arr3)")
	printArrayChange(arr3)
	fmt.Println("print arr3")
	for i, v := range arr3 {
		//下标和值
		fmt.Println(i, v)
	}

	//使用指针达到引用传递的效果
	fmt.Println("printArrayWithPointer(&arr3)")
	printArrayWithPointer(&arr3)
	fmt.Println("print arr3")
	for i, v := range arr3 {
		//下标和值
		fmt.Println(i, v)
	}
}
