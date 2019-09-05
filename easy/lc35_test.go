package easy

import (
	"fmt"
	"testing"
)

func Test35(t *testing.T) {
	nums := []int{1,3,5,6,7,9}
	val := 4
	b := searchInsert(nums,val)
	fmt.Println(b)
	fmt.Println(nums)
}

func searchInsert(nums []int, target int) int {
	index := 0

	for k,v := range nums {
		if v == target {
			index =k
			break
		}
		if v > target {
			tmp := append([]int{}, nums[k:]...)
			fmt.Println(tmp)
			nums = append(nums[:k], target)
			fmt.Println(nums)
			nums = append(nums, tmp...)
			fmt.Println(nums)
			index =k
			break
		}
	}
	return index
}