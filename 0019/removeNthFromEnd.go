package _019

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	front, rear := dummyHead, head
	for i := 0; i < n; i++ {
		rear = rear.Next
	}
	for rear != nil {
		front = front.Next
		rear = rear.Next
	}
	front.Next = front.Next.Next

	if front == dummyHead {
		head = head.Next
	}

	return head
}
