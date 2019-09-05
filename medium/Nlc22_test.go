package medium

import (
	"fmt"
	"testing"
)

func Test22(t *testing.T)  {
	a := 3
	s := generateParenthesis(a)
	fmt.Println(s)
}

func generateParenthesis(n int) []string {
	var s []string
	dfs(n,n,"",&s)
	return s
}
func dfs(left,right int, out string, s *[]string)  {

}