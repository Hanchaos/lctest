package easy

import (
	"fmt"
	"strconv"
	"testing"
)

func Test67(t *testing.T) {
	a:="11"
	b:="101"
	c := addBinary(a,b)
	fmt.Println(c)
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)

	var r string
	flag :=0
	for ; ;  {
		if la<=0 && lb<=0 {
			break
		}
		var sa, sb string
		if la<=0 {
			sa ="0"
		}else {
			sa = string(a[la-1])
		}
		if lb <= 0 {
			sb = "0"
		}else {
			sb = string(b[lb-1])
		}
		ia,_ := strconv.Atoi(sa)
		ib,_ := strconv.Atoi(sb)

		tmp := flag + ia + ib
		if tmp == 3{
			flag =1
			r = "1" +r
		}else if tmp ==2{
			flag =1
			r = "0" + r
		}else if tmp == 1{
			flag =0
			r = "1" +r
		}else {
			flag =0
			r = "0" +r
		}
		la --
		lb --
	}
	if flag == 1{
		r = "1" + r
	}

	return r
}