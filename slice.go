package main

import "fmt"

/**
一、Slice (切片) —— 一个范围，一个片段
	arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	s := arr[2:6]

二、Slice不是一个数据结构，是一种视图
	1.对Slice的操作实际上就是对arr的操作
	2.对Slice也能进行进一步的Slice操作
	3.Slice可以向后扩展，不可以向前扩展
	4.s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)

*/

func modifySlice(s []int) {
	s[0] = 100
}

func main() {
	//Slice的定义
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])
	fmt.Println()

	//证明Slice是一种视图
	s1 := arr[2:]
	fmt.Println("Before modifySlice(s1)")
	fmt.Println("s1 = ", s1)
	modifySlice(s1)
	fmt.Println("After modifySlice(s1)")
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println()

	s2 := arr[:6]
	fmt.Println("Before modifySlice(s2)")
	fmt.Println("s1 = ", s2)
	modifySlice(s2)
	fmt.Println("After modifySlice(s2)")
	fmt.Println(s2)
	fmt.Println(arr)
	fmt.Println()

	//不使用指针，使用Slice传参
	modifySlice(arr[:])
	fmt.Println(arr)
	fmt.Println()

	//Slice可以向后扩展，不可以向前扩展
	arr[0], arr[2] = 0, 2
	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Println("arr = ", arr)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Println(s3, s4)
	fmt.Println()
}
