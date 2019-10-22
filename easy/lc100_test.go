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

func Test107(t *testing.T) {

}
func levelOrderBottom(root *TreeNode) [][]int {
	var fuck map[int][]int
	fuck = make(map[int][]int)

	var a [][]int
	if root==nil{return a}

	f107(fuck,root,1)
	l := len(fuck)
	for i:=l;i>0;i-- {
		a = append(a,fuck[i] )
	}
	return a
}

func f107(fuck map[int][]int,p *TreeNode, dep int)  {
	if p == nil {
		return
	}
	fuck[dep] = append(fuck[dep], p.Val)
	if p.Left != nil {
		f107(fuck,p.Left,dep+1)
	}
	if p.Right !=nil {
		f107(fuck,p.Right,dep+1)
	}
}

func Test108(t *testing.T) {
	var a []int
	a = append(a, -10)
	a = append(a, -3)
	a = append(a, 0)
	a = append(a, 5)
	a = append(a, 9)
	sortedArrayToBST(a)
}
func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	root = fuck108(nums)
	return root
}

func fuck108(input []int) (t *TreeNode) {
	len1 := len(input)
	if 0 == len1 {
		return nil
	}
	middle := len1 / 2
	t = new(TreeNode)
	t.Val = input[middle]

	t.Left = fuck108(input[:middle])
	t.Right = fuck108(input[middle+1:])

	return t
}