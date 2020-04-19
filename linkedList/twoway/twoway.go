package main

import "fmt"

type Node struct {
	no   int
	data interface{}
	pre  *Node
	next *Node
}

type NodeOperator interface {
	InsertNode(headNode, newNode *Node)
	DeleteNode(headNode *Node, id int)
	SerializedListNode(headNode *Node)
	BackwardListNode(headNode *Node)
}

func (node *Node) InsertNode(headNode, newNode *Node) {
	temp := headNode
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.pre = temp
}

func (node *Node) DeleteNode(headNode *Node, id int) {
	temp := headNode
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		if temp.next.next != nil {
			temp.next = temp.next.next
			temp.next.pre = temp
		} else {
			temp.next = nil
		}
	} else {
		fmt.Printf("\n未发现节点:%d \n", id)
	}
}

// 正向输出
func (node *Node) SerializedListNode(headNode *Node) {
	temp := headNode
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d,%v====>]", temp.next.no, temp.next.data)
		// 判断是否有后续
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 反向输出
func (node *Node) BackwardListNode(headNode *Node) {
	temp := headNode
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d,%v====>]", temp.no, temp.data)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	head := &Node{}
	nodeOne := &Node{
		no:   0,
		data: map[string]string{"第一个节点": "想说点啥"},
	}
	nodeTwo := &Node{
		no:   1,
		data: "第二个节点",
	}
	nodeThree := &Node{
		no:   2,
		data: 1,
	}
	head.InsertNode(head, nodeOne)
	head.InsertNode(head, nodeTwo)
	head.InsertNode(head, nodeThree)
	head.SerializedListNode(head)
	head.DeleteNode(head, 2)
	head.DeleteNode(head, 4)
	fmt.Println()
	head.BackwardListNode(head)
}
