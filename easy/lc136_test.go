package easy

import (
	"fmt"
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