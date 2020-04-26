package main

import "fmt"

/**
重建二叉树
*/
type BinaryTreeNode struct {
	value     int
	leftNode  *BinaryTreeNode
	rightNode *BinaryTreeNode
}

func NewBinaryTree(value int) *BinaryTreeNode {
	return &BinaryTreeNode{
		value:     value,
		leftNode:  nil,
		rightNode: nil,
	}
}

func (node *BinaryTreeNode) PreOrder() {
	if node != nil {
		fmt.Printf("%d ", node.value)
		node.leftNode.PreOrder()
		node.rightNode.PreOrder()
	}
}

func (node *BinaryTreeNode) InOrder() {
	if node != nil {
		node.leftNode.InOrder()
		fmt.Printf("%d ", node.value)
		node.rightNode.InOrder()

	}
}

func CreateConstruct(preOrder []int, inOrder []int) *BinaryTreeNode {
	if preOrder == nil || inOrder == nil {
		return nil
	}
	return createConstruct(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

func createConstruct(preOrder []int, preStart int, preEnd int, inOrder []int, inStart, inEnd int) *BinaryTreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	root := NewBinaryTree(preOrder[preStart])
	var rootIndex int
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == root.value {
			rootIndex = i
			break
		}
	}
	leftCount := rootIndex - inStart
	root.leftNode = createConstruct(preOrder, preStart+1, preStart+1+leftCount, inOrder, inStart, rootIndex-1)
	root.rightNode = createConstruct(preOrder, preStart+leftCount+1, preEnd, inOrder, rootIndex+1, inEnd)
	return root
}
func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	bt := CreateConstruct(preOrder, inOrder)
	bt.PreOrder()
	fmt.Println()
	bt.InOrder()
}
