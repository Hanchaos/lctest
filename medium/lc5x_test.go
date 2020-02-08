package medium

import (
	"fmt"
	"testing"
)

func Test55(t *testing.T) {
	a:=[]int{2,3,1,1,4}
	fmt.Println(canJump(a))
	b:=make([]int, 5,10)
	fmt.Println(len(b),cap(b))
	b = append(b, 1,2,3,4,5)
	fmt.Println(len(b),cap(b))
}
func canJump(nums []int) bool {
	k:=0
	for i:=1;i<len(nums);i++{
		if i>k {return false}
		k = Max(k,i+nums[i])
	}
	return true
}

func Test56(t *testing.T) {
	a := []int{1,2,3}
	fmt.Println(subsets(a))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	tmp:=[]int{}
	helper(0,nums,tmp,&res)
	return res
}

func helper(n int, nums []int, tmp []int, res *[][]int)  {
	if len(tmp) <= len(nums){
		ll:=make([]int, len(tmp))
		copy(ll,tmp)
		fmt.Println(ll)
		*res= append(*res, ll)
	}
	for i:=n;i<len(nums);i++{
		tmp = append(tmp, nums[i])
		helper(i+1,nums,tmp,res)
		tmp = tmp[:len(tmp)-1]
	}
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	fuck94(root,&res)
	return res
}

func fuck94(node *TreeNode, res *[]int){
	if node==nil{
		return
	}
	fuck94(node.Left,res)
	*res= append(*res, node.Val)
	fuck94(node.Right,res)
}