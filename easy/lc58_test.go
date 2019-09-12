package easy

import (
	"fmt"
	"testing"
)

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
