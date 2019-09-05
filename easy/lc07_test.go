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
