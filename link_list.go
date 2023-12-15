package main

import "fmt"

type LinkedList struct {
	Head   *Node
	length int
}

type Node struct {
	Data int
	Next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (l *LinkedList) InsertAtTail(data int) {
	node := &Node{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
		l.length++

		return
	}

	curNode := l.Head
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = node

	l.length++
}

func (l *LinkedList) InsertAtHead(data int) {
	node := &Node{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
		l.length++

		return
	}

	node.Next = l.Head
	l.Head = node

	l.length++
}

func (l *LinkedList) InsertAt(data int, n int) {
	node := &Node{
		Data: data,
		Next: nil,
	}

	if n == 0 {
		l.InsertAtHead(data)

		return
	}
	if n >= l.length-1 {
		l.InsertAtTail(data)

		return
	}

	curNode := l.Head
	for i := 0; i < n-1; i++ {
		curNode = curNode.Next
	}
	node.Next = curNode.Next
	curNode.Next = node

	l.length++
}

func (l *LinkedList) Print() {
	if l.Head == nil {
		fmt.Println("Список пуст")

		return
	}

	curNode := l.Head
	for curNode != nil {
		fmt.Println(fmt.Sprintf("Value: %d, Next: %p", curNode.Data, curNode.Next))
		curNode = curNode.Next
	}
}

func (l *LinkedList) Get(position int) *Node {
	if l.Head == nil {
		fmt.Println("Список пуст")

		return nil
	}

	curNode := l.Head
	for i := 0; i < position; i++ {
		curNode = curNode.Next
	}

	return curNode
}

func PrintListNode(n *Node) {
	if n == nil {
		fmt.Println("Список пуст")

		return
	}
	fmt.Println(fmt.Sprintf("Value: %d, Next: %p", n.Data, n.Next))
}

//func main() {
//	list := NewLinkedList()
//	list.InsertAtTail(1)
//	list.InsertAtTail(2)
//	list.InsertAtTail(3)
//
//	list.InsertAtHead(0)
//	list.InsertAt(77, 2)
//
//	list.Print()
//
//	fmt.Println()
//
//	PrintListNode(list.Get(2))
//}
