package easy

import (
	"fmt"
	"testing"
	"time"
)

func Test53(t *testing.T) {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	b := maxSubArray(a)
	fmt.Println(b)
}

func maxSubArray(nums []int) int {
	return haha54(nums ,0,len(nums)-1)
}

func haha54(nums []int,left, right int) int {
	mid := (right+left)/2
	fmt.Println(left,right,mid)
	time.Sleep(1*time.Second)
	if left == right {
		return nums[left]
	}
	lans := haha54(nums,left,mid)
	rans := haha54(nums,mid+1,right)

	sum := 0
	lmax := nums[mid]
	for i:=mid;i>=left;i--{
		sum = sum+nums[i]
		if sum > lmax {
			lmax = sum
		}
	}
	sum = 0
	rmax := nums[mid+1]
	for i:=mid+1;i<=right;i++{
		sum = sum+nums[i]
		if sum > rmax {
			rmax = sum
		}
	}

	ans := rmax + lmax
	if lans > ans {
		ans = lans
	}
	if rans > ans{
		ans = rans
	}
	return ans
}

func Test58(t *testing.T) {
	a := "asasda asdas ss asd "
	//a = "a"
	b := lengthOfLastWord(a)
	fmt.Println(b)
}

func lengthOfLastWord(s string) int {
	n := len(s)
	flag := 0
	lengh := 0
	for i:=n-1;i>=0;i--{
		if string(s[i])!=" "{
			lengh++
			flag = 1
		}
		if flag ==1 && string(s[i])==" " {
			break
		}
	}

	return lengh
}