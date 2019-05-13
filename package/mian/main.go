package main

import "fmt"
import ".."

func main() {
	//创建对象的不同方法，不论是对象还是指针都用 "." 访问字段
	//var root TreeNode
	root := myPackage.TreeNode{Value: 3}
	root.Left = &myPackage.TreeNode{Value: 2}
	root.Right = &myPackage.TreeNode{5, nil, nil}
	root.Left.Right = new(myPackage.TreeNode)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
	fmt.Println()

	//使用工厂函数创建对象
	root.Right.Left = myPackage.CreateTreeNode(2)
	fmt.Println(root.Left)
	fmt.Println()

	//函数 即可以使用指针调用，也能用对象调用，很方便
	root.Print()

	root.SetValue(5)
	root.Print()

	pRoot := &root
	pRoot.SetValue(10)
	pRoot.Print()

	//使用nil指针调用方法
	var npRoot *myPackage.TreeNode
	npRoot.SetValue(100) //panic: runtime error: invalid memory address or nil pointer dereference

	//函数调用
	root.Traverse()
	fmt.Println()

	//扩展，后序遍历
	myTreeNode := myPackage.MyTreeNode{&root}
	myTreeNode.PostOrder()
}
