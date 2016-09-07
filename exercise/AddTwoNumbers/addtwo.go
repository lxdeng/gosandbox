package main

/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	test()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	head := new(ListNode)
	tail := head

	for l1 != nil && l2 != nil {
		value := l1.val + l2.val + carry
		carry = value / 10
		value = value % 10

		node := &ListNode{val: value}
		tail.next = node
		tail = node

		l1 = l1.next
		l2 = l2.next
	}

	l := l1
	if l2 != nil {
		l = l2
	}

	for l != nil {
		value := l.val + carry
		carry = value / 10
		value = value % 10

		node := &ListNode{val: value}
		tail.next = node
		tail = node

		l = l.next
	}

	if carry > 0 {
		tail.next = &ListNode{val: carry}
	}

	return head.next
}

func makeList(numbers []int) *ListNode {
	head := new(ListNode)
	tail := head

	for _, n := range numbers {
		node := &ListNode{val: n}
		tail.next = node
		tail = node
	}

	return head.next
}

func makeSlice(l *ListNode) []int {
	result := []int{}

	for l != nil {
		result = append(result, l.val)
		l = l.next
	}
	return result
}

func printList(l *ListNode) {
	if l == nil {
		fmt.Println("List: empty list")
	} else {
		fmt.Print("List: ")
		for l != nil {
			fmt.Printf("%d ", l.val)
			l = l.next
		}
		fmt.Println()
	}
}

func test() {
	nums1 := []int{2, 4, 3}
	l1 := makeList(nums1)
	printList(l1)

	nums2 := []int{5, 6, 4}
	l2 := makeList(nums2)
	printList(l2)

	r := addTwoNumbers(l1, l2)
	printList(r)
}
