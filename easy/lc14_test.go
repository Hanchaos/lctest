package easy

import (
	"fmt"
	"testing"
)

func Test14(t *testing.T)  {
	a := []string{"glower","flow","flowsss"}
	b := longestCommonPrefix(a)
	fmt.Println(b)
}

func longestCommonPrefix(strs []string) string {
	a := ""
	index := 0
	if len(strs)==0{
		return a
	}
	for {
		flag := 0
		if index>= len(strs[0]){
			break
		}
		tmp := string(strs[0][index])
		fmt.Println(tmp)
		for _, v := range strs {
			if index >= len(v){
				flag = 1
				break
			}
			if string(v[index]) != tmp{
				flag = 1
			}
		}
		if flag == 0 {
			a = a + tmp
		} else {
			break
		}
		index++
	}
	return a
}