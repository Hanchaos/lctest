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