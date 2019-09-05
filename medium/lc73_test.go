package medium

import (
	"fmt"
	"testing"
)
//矩阵置零
func Test73(t *testing.T) {
	in := [][]int {
		{1,4,3,4},
		{0,8,2,5},
		{8,4,6,6},
	}
	n:=len(in)
	m:=len(in[0])
	var row0 []int
	var col0 []int
	for i:=0;i<n;i++{
		fmt.Println(in[i])
		for j:=0;j<m;j++{
			if in[i][j] == 0{
				fmt.Println(i,j)
				row0 = append(row0, i)
				col0 = append(col0, j)
			}
		}
	}
	fmt.Println("----------------------------------------")
	for _,v := range row0{
		for j:=0;j<m;j++{
			in[v][j] = 0
		}
	}
	for i:=0;i<n;i++{
		fmt.Println(in[i])
	}
	fmt.Println("----------------------------------------")
	for _,v := range col0{
		for j:=0;j<n;j++{
			in[j][v] = 0
		}
	}
	for i:=0;i<n;i++{
		fmt.Println(in[i])
	}
}
