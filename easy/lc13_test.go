package easy

import (
	"fmt"
	"testing"
)

func Test13(t *testing.T)  {
	a := "MCMXCIV"
	b := romanToInt(a)
	fmt.Println(b)
}

func romanToInt(s string) int {
	x := 0
	n := len(s)
	m := 0
    for {
		if m == n {
			break
		}
		if m+1 != n {
			tmp := check(string(s[m]) + string(s[m+1]))
			if tmp != 0 {
				x = x + tmp
				m = m + 2
				continue
			}
		}
		x = x+check(string(s[m]))
    	m++
	}
	return x
}

func check(s string) int {
	switch s{
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	}

	return 0
}

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