package tree

//二叉树的操作
type Operator interface {
	PrintBinaryTree()
	Depth() int
	LeafCount() int
}

type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

