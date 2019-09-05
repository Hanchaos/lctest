package easy

import (
	"testing"
)

func Test09(t *testing.T)  {
	x := 12345432
	isPalindrome(x)
}

func isPalindrome(x int) bool {
	z := x
	y := 0
	for {
		tmp := x%10
		x = x/10
		y = y*10 + tmp
		if x == 0 {
			break
		}
	}
	if z < 0 {
		return false
	}else if y == z{
		return true
	}else {
		return false
	}
}