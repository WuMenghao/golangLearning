package myPackage

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//TreeNode 的工厂方法
func CreateTreeNode(Value int) *TreeNode {
	return &TreeNode{Value: Value}
}

//func 定义 有接收者
func (node *TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nie pointer ignored")
		return
	}
	node.Value = Value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
