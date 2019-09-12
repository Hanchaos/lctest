package easy

import (
	"fmt"
	"testing"
)

func Test66(t *testing.T) {
	a:= []int{9,9,9}
	b:=plusOne(a)
	fmt.Println(b)
}

func plusOne(digits []int) []int {
	l := len(digits)-1
	flag := 1
	for ; ;  {
		if l <0 {
			break
		}
		if digits[l] == 9 && flag == 1 {
			digits[l] = 0
			flag = 1
		}else if flag ==1 {
			digits[l]++
			flag = 0
		}else {
			flag = 0
		}
		l--
	}
	if flag ==1 {
		digits = append([]int{1},digits... )
	}
	return  digits
}