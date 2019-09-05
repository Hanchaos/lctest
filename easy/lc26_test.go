package easy

import (
	"fmt"
	"testing"
)

func Test26(t *testing.T)  {
	a := []int{0,0,1,1,2,3,4,4,4}
	b := removeDuplicates(a)
	fmt.Println(a)
	println(b)
}

func removeDuplicates(nums []int) int {
	index := 0
	if len(nums)==0{
		return 0
	}
	tmp := nums[0]-1
	for _,v := range nums{
		if tmp != v {
			nums[index]=v
			tmp = v
			index++
		}
	}
	return index
}