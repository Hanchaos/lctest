package medium

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T) {
	n:=3
	res:=make([]string,0)
	fuck22(&res,"",n,n)
	fmt.Println(res)
}

func fuck22(res *[]string, tmp string, left,right int) {
	if left==0&&right==0{
		*res = append(*res, tmp)
		return
	}
	if left>0{
		tmp = tmp + "("
		fuck22(res,tmp,left-1,right)
		tmp=tmp[:len(tmp)-1]
	}
	if right>0 && right>left{
		tmp=tmp+")"
		fuck22(res,tmp,left,right-1)
		tmp=tmp[:len(tmp)-1]
	}
}
