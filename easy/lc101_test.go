package easy

import "testing"

func Test101(t *testing.T) {

}

func isSymmetric(root *TreeNode) bool {
	return sy(root,root)
}

func sy(p *TreeNode,q *TreeNode) bool{
	if p==nil && q== nil{
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val{
		return false
	}
	if !sy(p.Right,q.Left){
		return false
	}
	if !sy(q.Right,p.Left){
		return false
	}

	return true
}
