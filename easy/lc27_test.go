package easy

import (
	"fmt"
	"testing"
)

func Test27(t *testing.T) {
	nums := []int{3,2,4,3,3,2,6,7,2,3}
	val := 3

	b := removeElement(nums,val)
	fmt.Println(nums)
	fmt.Println(b)
}

func removeElement(nums []int, val int) int {
	index := 0
	if len(nums)==0 {
		return 0
	}
	for _,v := range nums{
		if v != val {
			nums[index]=v
			index++
		}
	}
	return index
}