package main

import (
	"binaryTree/tree"
	"fmt"
)

func main() {
	//root := node.NewNode(nil, nil)
	//var initialTree node.Initer
	//initialTree = root
	//initialTree.SetData("rootNode")
	//
	//leftNode := node.NewNode(nil, nil)
	//leftNode.SetData("leftNode")
	//
	//secondLeftNode := node.NewNode(nil, nil)
	//secondRightNode := node.NewNode(nil, nil)
	//secondLeftNode.SetData(1000)
	//secondRightNode.SetData(3.141592653589)
	//leftNode.LeftChild = secondLeftNode
	//leftNode.RightChild = secondRightNode
	//
	//rightNode := node.NewNode(nil, nil)
	//rightNode.SetData("rightNode")
	//root.LeftChild = leftNode
	//root.RightChild = rightNode
	//root.PrintBinaryTree()
	//fmt.Println()
	//fmt.Printf("leaf: %d \n", root.LeafCount())
	//fmt.Printf("depth: %d\n", root.Depth())
	//var order tree.Order
	//order = root
	//fmt.Println("前序")
	//order.PreOrder()
	//fmt.Println()
	//fmt.Println("中序")
	//order.InOrder()
	//fmt.Println()
	//fmt.Println("后序")
	//order.PostOrder()
	bst := tree.BinarySearchTree{}
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
	bst.String()
	fmt.Println(bst.Max())
	fmt.Println(bst.Min())

}
