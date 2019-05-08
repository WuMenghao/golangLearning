package main

import "fmt"

/**
一、面向对象(oop)
	1.go语言仅支持封装，不支持继承和多态
	2.go语言没有class ，只有struct
	3.struct的创建是在堆上还是栈上，我们无需关心，自动GC

二、工厂函数
	使用自定义工厂函数，返回值为局部变量的地址
*/

//TreeNode 类型
type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (node TreeNode) print() {
	fmt.Println(node.value)
}

//TreeNode 的工厂方法
func createTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func main() {
	//创建对象的不同方法，不论是对象还是指针都用 "." 访问字段
	//var root TreeNode
	root := TreeNode{value: 3}
	root.left = &TreeNode{value: 2}
	root.right = &TreeNode{5, nil, nil}
	root.left.right = new(TreeNode)
	fmt.Println(root)
	fmt.Println(root.left)
	fmt.Println(root.right)
	fmt.Println()

	//使用工厂函数创建对象
	root.right.left = createTreeNode(2)
	fmt.Println(root.left)

	root.print()
	fmt.Println()
}
