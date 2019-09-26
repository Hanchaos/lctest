package easy

import (
	"fmt"
	"testing"
)

func Test88(t *testing.T) {
	a := []int{1,2,3,0,0,0}
	b := []int{2,5,6}
	merge(a,3,b,3)
	fmt.Println(a)
}

func merge( nums1 []int, m int, nums2 []int, n int)  {
	for m > 0 || n > 0 {
		if n == 0 {
			break
		}
		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}