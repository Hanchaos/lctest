package easy

import (
	"fmt"
	"sort"
	"testing"
)

func Test136(t *testing.T) {
	var a []int
	a = append(a, 4,1,2,1,2)
	singleNumber(a)
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println(res,nums[i])
		res ^= nums[i]
		fmt.Println(res)
	}
	return res
}

func Test617(t *testing.T) {

}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var root = t1
	root.Val = t1.Val + t2.Val
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}

func Test581(t *testing.T) {
	a:=[]int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(a))
}
func findUnsortedSubarray(nums []int) int {
	a:=make([]int, len(nums))
	copy(a,nums)
	sort.Ints(a)
	start:=0
	end:=0
	for i:=0;i<len(nums) ;i++  {
		if a[i]!=nums[i]{
			start=i
			break
		}
	}
	for i:=len(nums)-1;i>=0 ;i--  {
		if a[i]!=nums[i]{
			end=i
			break
		}
	}
	fmt.Println(start,end)
	return end-start+1
}

func Test560(t *testing.T) {
	a:=[]int{1,1,1}
	fmt.Println(subarraySum(a,2))
}
func subarraySum(nums []int, k int) int {
	res:=0
	l:=len(nums)
	mmap:=make(map[int]int)
	mmap[0]=1
	sum:=0
	for i:=0;i<l;i++{
		sum+=nums[i]
		res+=mmap[sum-k]
		mmap[sum]++
	}
	fmt.Println(sum)
	fmt.Println(mmap)
	return res
}