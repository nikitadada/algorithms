package main

import "fmt"

type LinkedList struct {
	Head   *ListNode
	length int
}

type ListNode struct {
	Data int
	Next *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (l *LinkedList) InsertAtTail(data int) {
	node := &ListNode{
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
	node := &ListNode{
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
	node := &ListNode{
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

func (l *LinkedList) Reverse() {
	var prev *ListNode = nil
	cur := l.Head

	for cur != nil {
		tmp := cur

		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	l.Head = prev
}

func (l *LinkedList) RemoveFromEnd(n int) {
	dummyNode := &ListNode{
		Next: l.Head,
	}
	curr := dummyNode
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	curr = dummyNode
	for i := 0; i < count-n-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
}

func (l *LinkedList) RemoveFromEnd2(n int) {
	dummyNode := &ListNode{
		Next: l.Head,
	}
	fast := dummyNode
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	slow := dummyNode
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	l.Head = dummyNode.Next
}

func (l *LinkedList) MiddleNode() *ListNode {
	slow := l.Head
	fast := l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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

func (l *LinkedList) Get(position int) *ListNode {
	if l.Head == nil {
		return nil
	}

	curNode := l.Head
	for i := 0; i < position; i++ {
		curNode = curNode.Next
	}

	return curNode
}

func main() {
	list := NewLinkedList()
	list.InsertAtTail(1)
	list.InsertAtTail(2)
	list.InsertAtTail(3)
	list.InsertAtTail(4)
	list.InsertAtTail(5)

	list.Print()
	fmt.Println()

	PrintListNode(list.MiddleNode())
}

func PrintListNode(n *ListNode) {
	if n == nil {
		fmt.Println("Список пуст")

		return
	}
	fmt.Println(fmt.Sprintf("Value: %d, Next: %p", n.Data, n.Next))
}
