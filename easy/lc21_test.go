package easy

import "testing"

func Test21(t *testing.T){

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode
	result = new(ListNode)

	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	p1 := l1
	p2 := l2
	tmp := result

	for {
		if p1==nil || p2==nil{
			break
		}
		if p1.Val <= p2.Val {
			tmp.Next = p1
			tmp = tmp.Next
			p1 = p1.Next
		}else {
			tmp.Next = p2
			tmp = tmp.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		tmp.Next = p1
	}
	if p2 != nil {
		tmp.Next = p2
	}

	return result.Next // result self has 0
}

type ListNode struct {
	Val int
	Next *ListNode
}

