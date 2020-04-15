package dataStruct

import "fmt"

type LinkedListNode struct {
	Next *LinkedListNode
	Val  interface{}
}

func (l *LinkedListNode) AddLinkedListNode(tail *LinkedListNode, val interface{}) *LinkedListNode {

	tail.Next = &LinkedListNode{Val: val}

	return tail.Next

}

// 지우고자 하는 노드, root, tail 노드 입력 받음
func RemoveLinkedListNode(node *LinkedListNode, root *LinkedListNode, tail *LinkedListNode) (*LinkedListNode, *LinkedListNode) {

	prev := root

	for prev.Next != node {
		prev = prev.Next
	}

	// 지우고자 하는 노드가 root 면
	if root == node {
		root = node
		node.Next = nil

	} else if tail == node { // 지우고자 하는 노드가 tail 이면
		prev.Next = nil
	} else {
		prev.Next = prev.Next.Next
	}
	return root, tail

}

func PrintLinkedList(root *LinkedListNode) {
	for root.Next != nil {
		fmt.Printf("%v-> ", root.Val)
		root = root.Next
	}
	fmt.Printf("%v", root.Val)
}
