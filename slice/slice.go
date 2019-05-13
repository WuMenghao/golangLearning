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

三、Slice的操作
1）append 操作 s := append(s,val)
	1.添加元素时如果超远cap,系统会重新分配更大的底层数组
	2.由于值传递的关系，必须接受append的返回值
2)create 操作 var s []int
	1.slice的初始值是nil
*/

func modifySlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("%v len=%d ,cap=%d\n", s, len(s), cap(s))
}

func part01() {
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

	//向Slice添加元素 s4 s5不再是对arr2的view了，系统重新分配更大的数组
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr2[2:6]
	s2 = s1[3:5]
	fmt.Println("arr2 = ", arr2)
	s3 = append(s2, 10)
	s4 = append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	fmt.Println("s3 = ", s3)
	fmt.Println("s4 = ", s4)
	fmt.Println("s5 = ", s5)
	fmt.Println("arr2 = ", arr2)
	fmt.Println()
}

func part02() {
	//创建Slice
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	//创建Slice初始化
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//创建已知长度的Slice
	s2 := make([]int, 16)
	printSlice(s2)

	//创建已知len和cap的Slice
	s3 := make([]int, 10, 32)
	printSlice(s3)

	//切片拷贝
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	//删除元素
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	//删除前后元素
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}

func main() {
	//part01()
	part02()
}
