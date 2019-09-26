package easy

import "testing"

func Test104(t *testing.T) {
	
}

func maxDepth(root *TreeNode) int {
	return f104(root,0)
}

func f104(p *TreeNode, dep int) int{
	ml :=0
	mr :=0
	if p==nil{
		return 0
	}
	if p.Left ==nil && p.Right==nil{
		return dep
	}
	if p.Left !=nil {
		ml = f104(p.Left,dep+1)
	}
	if p.Right !=nil {
		mr = f104(p.Right,dep+1)
	}
	if mr > ml {
		return mr
	}else {
		return ml
	}
}