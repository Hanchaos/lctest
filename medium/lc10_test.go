package medium

import (
	"fmt"
	"sort"
	"testing"
)

func Test11(t *testing.T) {
	var a = []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func maxArea(height []int) int {
	l:=0
	r:=len(height)-1
	max:=0;
	for ; ;  {
		if l>=r{ break }
		v:= Min(height[l], height[r]) * (r-l)
		max=Max(max,v)
		if height[l]<=height[r] {
			l++
		}else {
			r--
		}
	}
	return max
}

func Test15(t *testing.T) {
	var a = []int{-1, 0, 1, 2, -1, -4}
	b:=threeSum(a)
	fmt.Println(b)
}
func threeSum(nums []int) [][]int {
	res := [][]int{}
	ll:=len(nums)
	if len(nums)==0 {return res}
	sort.Ints(nums)
	fmt.Println(nums)
	for i:=0; i<ll-2;i++  {
		l:=i+1
		r:=ll-1
		for l<r  {
			if nums[l]+nums[r]== -nums[i]{
				fmt.Println(i,l,r)
				tmp := []int{}
				tmp = append(tmp, nums[l],nums[i],nums[r])
				res = append(res, tmp)
				t:=nums[l]
				for t==nums[l] {
					fmt.Println(l)
					l++}
			}else if nums[l]+nums[r]<nums[i]{
				l++
			}else {
				r--
			}
		}
		for nums[i]==nums[i+1]{
			i++
		}
	}
	return res
}

func Test17(t *testing.T) {
	a := ""
	b:=letterCombinations(a)
	fmt.Println(b)
}
func letterCombinations(digits string) []string {
	res:= make([]string,0)
	if digits=="" {return res}
	m := map[string][]string{
		"2":{"a","b","c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res = append(res, "")
	for _,v := range digits{
		k:=string(v)
		pre:=res
		res = []string{}
		fmt.Println(m[k])
		for _, num := range m[k]{

			for _, pv := range pre {
				res = append(res, pv+num)
			}
		}
	}
	return res
}

func Test19(t *testing.T) {
	
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for n>0 {
		fast=fast.Next
		n--
	}
	for fast != nil {
		fast=fast.Next
		slow=slow.Next
	}
	return slow

}

func Test111(t *testing.T) {
	a:=[]int{2,3,4,5,1}
	fuckkk(a)
	b:=5
	shittt(b)
	fmt.Println(a,b)
}
func fuckkk(a []int){
	sort.Ints(a)
}
func shittt(b int)  {
	b=10
}