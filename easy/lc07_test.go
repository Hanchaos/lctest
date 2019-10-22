package easy

import (
	"fmt"
	"testing"
)

func Test07(t *testing.T)  {
	a := -123456780
	b := 0
	for {
		c:=a%10
		a = a/10
		b = b*10+c
		if a == 0 {
			break
		}
	}
	if b> 2147483647 || b < -2147483648 {
		fmt.Println(0)
	}
	fmt.Println(b)
}
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