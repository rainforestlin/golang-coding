package node

import "fmt"

type Node struct {
	Data       interface{}
	LeftChild  *Node
	RightChild *Node
}

//创建新节点的接口
type Initer interface {
	SetData(data interface{})
}

// 创建新节点
func NewNode(left, right *Node) *Node {
	return &Node{
		Data:       nil,
		LeftChild:  left,
		RightChild: right,
	}
}

// 赋值给节点
func (node *Node) SetData(data interface{}) {
	node.Data = data
}

// 打印二叉树
func (node *Node) PrintBinaryTree() {
	if node != nil {
		fmt.Printf("%v ", node.Data)
		if node.LeftChild != nil || node.RightChild != nil {
			fmt.Printf("(")
			node.LeftChild.PrintBinaryTree()
			if node.RightChild != nil {
				fmt.Printf(",")
			}
			node.RightChild.PrintBinaryTree()
			fmt.Printf(")")
		}
	}
}

// 二叉树深度
func (node *Node) Depth() int {
	var depthLeft, depthRight int
	if node == nil {
		return 0
	} else {
		depthLeft = node.LeftChild.Depth()
		depthRight = node.RightChild.Depth()
		if depthLeft > depthRight {
			return depthLeft + 1
		} else {
			return depthRight + 1
		}
	}
}

// 叶子节点数
func (node *Node) LeafCount() int {
	if node == nil {
		return 0
	} else if (node.LeftChild == nil) && (node.RightChild == nil) {
		return 1
	} else {
		return node.RightChild.LeafCount() + node.LeftChild.LeafCount()
	}
}

// 前序遍历
func (node *Node) PreOrder() {
	if node != nil {
		fmt.Printf("%v ", node.Data)
		node.LeftChild.PreOrder()
		node.RightChild.PreOrder()
	}
}

// 中序遍历
func (node *Node) InOrder() {
	if node != nil {
		node.LeftChild.InOrder()
		fmt.Printf("%v ", node.Data)
		node.RightChild.InOrder()
	}
}

// 后序遍历
func (node *Node) PostOrder() {
	if node != nil {
		node.LeftChild.PostOrder()
		node.RightChild.PostOrder()
		fmt.Printf("%v ", node.Data)
	}
}
