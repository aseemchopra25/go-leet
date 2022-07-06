package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := new(ListNode)
	sumHead := curr
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = new(ListNode)
		curr.Next.Val = sum % 10
		fmt.Println(curr.Next.Val)

		curr = curr.Next
		if sum > 9 {
			sum = 1
		} else {
			sum = 0
		}

	}
	if sum != 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = sum
	}
	return sumHead.Next
}

func main() {
	l1 := ListNode{}
	l2 := ListNode{}
	l3 := ListNode{}
	l1.Next = &l2
	l2.Next = &l3
	l1.Val = 2
	l2.Val = 4
	l3.Val = 3

	l4 := ListNode{}
	l5 := ListNode{}
	l6 := ListNode{}
	l4.Next = &l5
	l5.Next = &l6
	l4.Val = 5
	l5.Val = 6
	l6.Val = 4

	x := addTwoNumbers(&l1, &l4)

	fmt.Println(x.Val, x.Next.Val, x.Next.Next.Val)

}
