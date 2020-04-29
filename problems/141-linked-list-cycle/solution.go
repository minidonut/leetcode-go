package main

import "github.com/minidonut/leetcode-go/pkg/containers"

type ListNode = containers.ListNodeIntCyclic

func Solve(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
