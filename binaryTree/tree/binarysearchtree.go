package tree

import (
	"fmt"
	"sync"
)

type BinarySearchTreeOperator interface {
	Insert(no int, value interface{}) // 向二叉搜索树的合适位置插入节点
	Search(no int) bool               // 检查序号为no值的元素是否在树中存在
	Remove(key int)                   // 删除所有值位 value 的节点
	Min() interface{}                 // 获取二叉搜索树中的最小值
	Max() interface{}                 // 获取二叉搜索树中的最大值
	Order                             // 遍历方式
	String()                          // 打印出二叉树
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

type Node struct {
	key   int
	value interface{}
	left  *Node //left
	right *Node //right
}

func (bst *BinarySearchTree) Insert(no int, value interface{}) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{
		key:   no,
		value: value,
		left:  nil,
		right: nil,
	}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// 中序遍历
func (bst *BinarySearchTree) InOrder() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	if bst.root != nil {
		inOrder(bst.root)
	}
}

// 前序遍历
func (bst *BinarySearchTree) PreOrder() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	if bst.root != nil {
		preOrder(bst.root)
	}

}

func (bst *BinarySearchTree) PostOrder() {
	bst.lock.RLock()
	bst.lock.RUnlock()
	if bst.root != nil {
		postOrder(bst.root)
	}

}

func (bst *BinarySearchTree) Min() interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	node := bst.root
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return (*node).value
		}
		node = node.left
	}
}

func (bst *BinarySearchTree) Max() interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	node := bst.root
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
			return (*node).value
		}
		node = node.right
	}
}

func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

func search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}
	return true
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	remove(bst.root, key)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return node
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	findSmallestRight := node.right
	for {
		if findSmallestRight != nil && findSmallestRight.left != nil {
			findSmallestRight = findSmallestRight.left
		} else {
			break
		}
	}
	node.key, node.value = findSmallestRight.key, findSmallestRight.value
	node.right = remove(node.right, node.key)
	return node
}

func (bst *BinarySearchTree) String() {
	bst.lock.RLock()
	bst.lock.RUnlock()
	fmt.Println("------------------------------------------------")
	stringfy(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringfy(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringfy(node.right, level)
		fmt.Printf(format+"%d\n", node.key)
		stringfy(node.left, level)
	}
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%v ", node.value)
		inOrder(node.right)
	}
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%v ", node.value)
		preOrder(node.left)
		preOrder(node.right)
	}
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Printf("%v ", node.value)
	}
}
