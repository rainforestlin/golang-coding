package main

import "fmt"

// 声明一个单向链表的节点
type Node struct {
	no   int
	data interface{}
	next *Node
}

type NodeOperator interface {
	InsertNode(headNode, newNode *Node)
	ListNode(headNode *Node)
	DeleteNode(headNode *Node, id int)
}

func (node *Node) InsertNode(headNode, newNode *Node) {
	temp := headNode
	for {
		// 最后一个节点
		if temp.next == nil {
			break
		}
		// 让temp不断指向下一个节点
		temp = temp.next
	}
	temp.next = newNode
}

func (node *Node) ListNode(headNode *Node) {
	// 创建一个临时节点
	temp := headNode
	// 先判断是否为一个空链表
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

// 删除节点
func (node *Node) DeleteNode(headNode *Node, id int) {
	temp := headNode
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			// 找到节点
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("未找到")
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
	head.ListNode(head)
}
