package easy

import (
	"fmt"
	"strconv"
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

/*
1.     1
2.     11
3.     21
4.     1211
5.     111221*/
func Test38(t *testing.T) {
	fuck := 5
	shit := countAndSay(fuck)
	fmt.Println(shit)
}

func countAndSay(n int) string {
	if n == 1{
		return "1"
	}
	result := "1"
	for begin:= 1; begin<n; begin++ {
		result = haha38(result)
	}
	return result
}


func haha38(s string) string {
	result := ""
	tmp := ""
	num := 0
	for k,v := range s {
		if k == 0 {
			tmp = string(v)
			num = 1
		}else if tmp == string(v) {
			num++
		}else {
			result = result + strconv.Itoa(num) + tmp
			tmp = string(v)
			num = 1
		}
	}
	if num > 0 {
		result = result + strconv.Itoa(num) + tmp
	}
	return result
}