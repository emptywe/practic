package main

import (
	"fmt"
)

type ListNode struct {
	Next *ListNode
	val  int
}

func NewList(val int) *ListNode {
	return &ListNode{val: val}
}

func add(list *ListNode, val int) *ListNode {
	next := list
	list = &ListNode{val: val, Next: next}
	return list
}

func (l *ListNode) print() {
	list := l
	for list != nil {
		fmt.Print(list.val, " ")
		list = list.Next
	}
	fmt.Println()
}

func reverseList(listHead *ListNode) *ListNode {
	if listHead == nil {
		return nil
	}
	if listHead.Next == nil {
		return listHead
	}
	var prev, next *ListNode
	list := listHead
	for list.Next != nil {
		next = list.Next
		list.Next = prev
		prev = list
		list = next
	}
	list.Next = prev
	return list
}

func main() {

	List := NewList(1)
	for i := 2; i <= 10000; i++ {
		List = add(List, i)
	}
	List.print()

	reverseList(List).print()

}
