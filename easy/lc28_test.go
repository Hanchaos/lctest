package easy

import (
	"fmt"
	"testing"
)

func Test28(t *testing.T) {
	h := "hello"
	n := "ll"

	r := strStr(h,n)
	fmt.Println(r)
}

func strStr(haystack string, needle string) int {
	index := -1
	if needle == "" {
		return 0
	}

	for k,v := range haystack {
		if k+len(needle) >len(haystack){
			break
		}
		if string(v) == string(needle[0]) {
			index = k
			for m,n := range needle {
				if string(haystack[k+m]) != string(n) {
					index = -1
					break
				}
			}
		}
		if index != -1{
			break
		}
	}
	return index
}