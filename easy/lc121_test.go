package easy

import (
	"fmt"
	"strings"
	"testing"
)

func Test121(t *testing.T) {
	var a []int
	a = append(a, 7,1,5,3,6,4)
	b := maxProfit(a)
	fmt.Println(b)
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l==0{
		return 0
	}
	var a [][]int
	for i:=0;i<l;i++ {
		if i == 0 {
			var tmp []int
			tmp = append(tmp,0, -prices[i])
			a= append(a,tmp)
			continue
		}
		var tmp []int
		x := max(a[i-1][0],a[i-1][1]+prices[i]) // sell
		y := max(a[i-1][1],-prices[i]) // buy
		tmp = append(tmp,x,y)
		a = append(a, tmp)
	}
	return a[l-1][0]
}

func maxProfit122(prices []int) int {
	l := len(prices)
	if l==0{
		return 0
	}
	var a [][]int
	for i:=0;i<l;i++ {
		if i == 0 {
			var tmp []int
			tmp = append(tmp,0, -prices[i])
			a= append(a,tmp)
			continue
		}
		var tmp []int
		x := max(a[i-1][0],a[i-1][1]+prices[i]) // sell
		y := max(a[i-1][1],a[i-1][0]-prices[i]) // buy
		tmp = append(tmp,x,y)
		a = append(a, tmp)
	}
	return a[l-1][0]
}

func Test125(t *testing.T) {
	a :=" "
	isPalindrome125(a)
}

func isPalindrome125(s string) bool {
	if s == ""{
		return true
	}
	x := strings.ToLower(s)
	start :=0
	end := len(x)-1
	for ; ;  {
		v := x[start]
		w := x[end]
		if start >= end{
			break
		}
		if v<48 || (v >57 && v <97)|| v>122 {
			start ++
			continue
		}
		if w<48 || (w >57 && w <97)|| w>122 {
			end--
			continue
		}

		if v != w{
			return false
		}
		start ++
		end--

	}


	return true
}