package easy

import "testing"
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test100(t *testing.T) {



}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	flag := true
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val{
		if !isSameTree(p.Left,q.Left){
			flag=false
		}
		if !isSameTree(p.Right,q.Right){
			flag=false
		}
	}else {
		flag=false
	}

	return flag
}
