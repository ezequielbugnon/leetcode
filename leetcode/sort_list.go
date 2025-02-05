package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func getMiddle(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow := head
	fast := head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	return head, slow
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := getMiddle(head)
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return dummy.Next
}
