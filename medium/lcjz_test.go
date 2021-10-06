package medium

import (
	"fmt"
	"testing"
)

func TestJS(t *testing.T) {
	a:="abcabcbb"
	fmt.Println(lengthOfLongestSubstring(a))
}

func lengthOfLongestSubstring(s string) int {
	left:=0
	right:=0
	l:=len(s)
	a := make(map[byte]int)
	res:=0
	for right<l{
		c := s[right]
		a[c]++
		for a[c]>1 {
			c2 := s[left]
			a[c2]--
			left++
		}
		res = Max(res,right-left+1)
		right++
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	if root==nil{
		return res
	}
	helper102(root,&res,0)
	return res
}
func helper102(n *TreeNode, res *[][]int, lv int)  {
	if lv>=len(*res){
		tmp:=make([]int,0)
		*res = append(*res, tmp)
	}
	(*res)[lv] = append((*res)[lv], n.Val)
	if n.Right != nil { helper102(n.Right, res, lv+1)}
	if n.Left !=nil { helper102(n.Left,res,lv+1) }
}