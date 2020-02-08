package medium

import "testing"

type ListNode struct {
	Val int
	Next *ListNode
}

func Test02(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	addNode(l1,l2,0)
}

func addNode(l1 *ListNode, l2 *ListNode, pre int) (*ListNode){
	if l1==nil && l2==nil {
		if pre !=0 {
			return &ListNode{Val:pre}
		} else {
			return nil
		}
	}
	if l1 == nil{
		l1 = &ListNode{Val:0}
	}
	if l2 == nil{
		l2 = &ListNode{Val:0}
	}
	totalSum := l1.Val + l2.Val + pre
	i,j := totalSum / 10 , totalSum %10
	l := new(ListNode)
	l.Val = j
	l.Next = addNode(l1.Next, l2.Next, i)
	return l
}
