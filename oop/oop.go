package mian

import "fmt"

/**
一、面向对象(oop)
	1.go语言仅支持封装，不支持继承和多态
	2.go语言没有class ，只有struct
	3.struct的创建是在堆上还是栈上，我们无需关心，自动GC

二、工厂函数
	使用自定义工厂函数，返回值为局部变量的地址

三、方法接收者
	1.golang 的 struct中无法定义方法
	2.golang 中 func的定义可以指定接收者
	3.只有使用指针作为方法的接收者才可以改变结构内容
	4.nil指针也可以调用方法

四、值接收者和指针接收者
	1.要改变内容必须适应指针接收者
	2.结构过大也要考虑使用指针接收者
	3.一致性，如果有使用指针接收者，最好都使用指针接收者
	4.都是使用"."调用
	5.值接收者是golang特有的,java是指针接收者(引用传递)
*/

//TreeNode 类型
type TreeNode struct {
	value       int
	left, right *TreeNode
}

//TreeNode 的工厂方法
func createTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

//func 定义 有接收者
func (node *TreeNode) print() {
	fmt.Println(node.value)
}

func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nie pointer ignored")
		return
	}
	node.value = value
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
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
	fmt.Println()

	//函数 即可以使用指针调用，也能用对象调用，很方便
	root.print()

	root.setValue(5)
	root.print()

	pRoot := &root
	pRoot.setValue(10)
	pRoot.print()

	//使用nil指针调用方法
	var npRoot *TreeNode
	npRoot.setValue(100) //panic: runtime error: invalid memory address or nil pointer dereference

	//函数调用
	root.traverse()
}
