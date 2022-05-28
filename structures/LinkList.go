package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Ints(l *ListNode) []int {
	res := []int{}
	limit := 100
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
		limit--
		if limit < 0 {
			panic("reach the limit")
		}
	}
	return res
}

func Ints2List(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := &ListNode{Val: 0}
	current := head
	for _, val := range nums {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head.Next
}

func (l *ListNode) GetNodeWith(val int) *ListNode {
	for l != nil {
		if l.Val == val {
			return l
		}
		l = l.Next
	}
	return l
}

func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos < 0 {
		return head
	}
	cNode := head
	for pos > 0 {
		cNode = cNode.Next
		pos--
	}
	tail := cNode

	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = cNode

	return head
}
