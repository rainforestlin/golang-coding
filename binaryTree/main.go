package main

import (
	"binaryTree/node"
	"binaryTree/tree"
	"fmt"
)

func main() {
	root := node.NewNode(nil, nil)
	var initialTree node.Initer
	initialTree = root
	initialTree.SetData("rootNode")

	leftNode := node.NewNode(nil, nil)
	leftNode.SetData("leftNode")

	secondLeftNode := node.NewNode(nil, nil)
	secondRightNode := node.NewNode(nil, nil)
	secondLeftNode.SetData(1000)
	secondRightNode.SetData(3.141592653589)
	leftNode.LeftChild = secondLeftNode
	leftNode.RightChild = secondRightNode

	rightNode := node.NewNode(nil, nil)
	rightNode.SetData("rightNode")
	root.LeftChild = leftNode
	root.RightChild = rightNode
	root.PrintBinaryTree()
	fmt.Println()
	fmt.Printf("leaf: %d \n", root.LeafCount())
	fmt.Printf("depth: %d\n", root.Depth())
	var order tree.Order
	order = root
	fmt.Println("前序")
	order.PreOrder()
	fmt.Println()
	fmt.Println("中序")
	order.InOrder()
	fmt.Println()
	fmt.Println("后序")
	order.PostOrder()
}
