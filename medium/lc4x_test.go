package medium

import (
	"fmt"
	"testing"
)

func Test46(t *testing.T) {
	a:=[]int{1,2,3}
	b:=permute(a)
	fmt.Println(b)
}

func permute(nums []int) [][]int {
	r := [][]int{}
	S(nums, []int{}, &r)
	return r
}

func S(nums, t []int, r *[][]int) {
	l := len(nums)

	if l == 0 {
		*r = append(*r, t)
		return
	}

	for i := 0; i < l; i++ {
		a := make([]int, l)
		copy(a, nums)

		c := make([]int, len(t))
		copy(c, t)

		c = append(c, a[i])
		S(append(a[:i], a[i + 1:]...), c, r)
	}
}
