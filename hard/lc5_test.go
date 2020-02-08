package hard

import (
	"fmt"
	"testing"
)

func Test05(t *testing.T) {
	var b string
	b="ccccccccccccccccccccccccccc"
	a := longestPalindrome(b)
	fmt.Println(a)
}
func longestPalindrome(s string) string {
	if s==""{return  s}
	m:= make(map[int]map[int]bool)
	l:=len(s)
	b,e:=0,0
	for i:=0;i<l;i++ {
		im := make(map[int]bool)
		m[i] = im
		m[i][i]=true
	}
	for i:=1;i<l;i++ {
		for j:=0; j<i; j++ {
			if s[i]==s[j] && (i-j==1|| m[j+1][i-1]==true) {
				m[j][i]=true
				if i-j>e-b{
					b=j;e=i
				}
				continue
			}else {
				m[j][i]=false
			}
		}
	}
	fmt.Println(b,e)
	return s[b:e+1]
}