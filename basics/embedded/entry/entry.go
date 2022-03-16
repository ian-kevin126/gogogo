package main

import (
	"fmt"
	"gogogo/basics/tree"
)

// MyTreeNode 扩展已有的类型
type myTreeNode struct {
	//node *tree.Node
	*tree.Node // 在这一行去掉node，实现内嵌（embedding），是一个语法糖
}

// 扩展二叉树的后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}

	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse() // 0 2 3 4 5

	fmt.Println()
	root.postOrder() // 2 0 4 5 3
	fmt.Println()
}
