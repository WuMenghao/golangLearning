package myPackage

type MyTreeNode struct {
	Node *TreeNode
}

//扩展后序遍历
func (node *MyTreeNode) PostOrder() {
	if node == nil || node.Node == nil {
		return
	}
	left := MyTreeNode{node.Node.Left}
	left.PostOrder()
	right := MyTreeNode{node.Node.Right}
	right.PostOrder()
	node.Node.Print()
}
