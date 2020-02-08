package medium

import (
	"fmt"
	"sort"
	"testing"
)

func Test31(t *testing.T) {
	a:=[]int{1,4,3}
	nextPermutation(a)
	fmt.Println(a)
}
func nextPermutation(nums []int)  {
	length := len(nums)
	if length == 1 {
		return
	}
	var flag bool
	var i, k int

	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i + 1] { // 步骤1
			for k = length - 1; k > i; k-- { // 步骤2
				if nums[k] > nums[i] {
					break
				}
			}
			temp := nums[i]
			nums[i] = nums[k]
			nums[k] = temp
			flag = true
			break
		}
	}
	if flag { // 步骤3
		for k = i + 1; k <= ((i + 1 + length - 1) / 2); k++ {
			temp := nums[k]
			nums[k] = nums[length - k + i]
			nums[length - k + i] = temp
		}
	}else { // 步骤4
		for k = 0; k < (length / 2); k++ {
			temp := nums[k]
			nums[k] = nums[length - k - 1]
			nums[length - k - 1] = temp
		}
	}

	return
}

func Test33(t *testing.T) {

}

func Test34(t *testing.T) {
	a:=[]int{5,7,7,8,8,10}
	fmt.Println(searchRange(a,8))
}

func searchRange(nums []int, target int) []int {
	res := []int{-1,-1}
	l:=0
	r:=len(nums)-1
	for l<=r {
		fmt.Println(l,r)
		mid:=(l+r)/2
		if nums[mid]>target {
			r=mid-1
		}else if nums[mid]<target{
			l=mid+1
		}else {
			s,e := mid,mid
			fmt.Println(mid)
			for ; s != 0 && nums[s-1] == target; {s--}
			for ; e != len(nums)-1 && nums[e+1] == target; {e++}
			return []int{s,e}
		}
	}
	return res
}

func Test39(t *testing.T) {
	a:=[]int{2,3,5}
	b := combinationSum(a,8)
	fmt.Println(b)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res:=[][]int{}
	tmp:=[]int{}
	fuck39(candidates, &res, tmp, target, 0)
	return res
}

func fuck39(candidates []int, res *[][]int, tmp[]int, t int, num int)  {
	if t==0 {
		d := make([]int, len(tmp))
		copy(d, tmp)
		*res = append(*res, d)
		fmt.Println("OK",tmp)
		fmt.Println(*res)
		return
	}
	for i:=num; i<len(candidates);i++  {
		if i>0 && candidates[i]==candidates[i-1]{
			continue
		}
		v:=candidates[i]
		if t>=v {
			tmp = append(tmp, v)
			fuck39(candidates, res, tmp, t-v, i)
			tmp = tmp[:len(tmp)-1]
			continue
		}else {
			return
		}
	}
}