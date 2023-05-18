package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList - creates new linked list
func NewList(val int) *ListNode {
	return &ListNode{Val: val}
}

// Add - adds new list node to the end
func (l *ListNode) Add(val int) {
	list := l
	for list.Next != nil {
		list = list.Next
	}
	list.Next = &ListNode{Val: val}
}

// print - prints linked list
func (l *ListNode) print() {
	list := l
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
}

// length - returns list length
func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	if head.Next == nil {
		return 1
	}
	var l int

	for head != nil {
		head = head.Next
		l++
	}
	return l
}

// rotateRight - function from rotation task. rotate list by given k
func rotateRight(head *ListNode, k int) *ListNode {

	l := length(head)
	if l == 0 || l == 1 {
		return head
	}

	var (
		list     *ListNode
		prevList *ListNode
		iterator int
	)

	iterator = k % l

	for i := 0; i < iterator; i++ {
		list = head
		for list.Next != nil {
			prevList = list
			list = list.Next
		}
		list.Next = head
		prevList.Next = nil
		head = list
	}

	return head
}

// listArr - creates array with list node links
func listArr(head *ListNode) []*ListNode {
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	return arr
}

// reorderList - functions from reorder list task. Reorder list L1...Ln by formula: L0,Ln,L1,Ln-1...
func reorderList(head *ListNode) {
	listArray := listArr(head)
	counter := 0
	for i := 0; i < len(listArray)/2; i++ {
		listArray[i].Next = listArray[(len(listArray) - 1 - counter)]
		listArray[i].Next.Next = listArray[i+1]
		listArray[i+1].Next = nil
		counter++
	}
}

func main() {

	ListHead := NewList(1)
	ListHead.Add(2)
	ListHead.Add(3)
	ListHead.Add(4)
	ListHead.Add(5)
	ListHead.print()

	ListHead = rotateRight(ListHead, 1)

	fmt.Println()
	ListHead.print()

	reorderList(ListHead)

	fmt.Println()
	ListHead.print()

}
