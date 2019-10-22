package easy

import (
	"math"
	"testing"
	"fmt"
)

func Test110(t *testing.T) {
	
}

func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	l := fuck110(root.Left,0)
	r := fuck110(root.Right,0)
	if math.Abs(float64(r-l)) <=1  && isBalanced(root.Left) && isBalanced(root.Right){
		return true
	}
	return false
}
func fuck110(node *TreeNode, len int) int {
	if node == nil {
		return len
	}
	l := fuck110(node.Left,len+1)
	r := fuck110(node.Right,len+1)

	if r>l {
		return r
	}else {
		return l
	}

}

func Test111(t *testing.T) {

}

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Right !=nil && root.Left !=nil{
		return min(fuck111(root.Right,1),fuck111(root.Left,1))
	}
	return max(fuck111(root.Right,1),fuck111(root.Left,1))

}
func fuck111(node *TreeNode, len int) int {
	if node==nil{
		return 0
	}
	if node.Right !=nil && node.Left !=nil{
		return min(fuck111(node.Right,len),fuck111(node.Left,len))+1
	}
	return max(fuck111(node.Right,len),fuck111(node.Left,len))+1
}

func min(a,b int) int {
	if a>b{return b}
	return a
}
func max(a,b int) int {
	if a>b{return a}
	return b
}

func Test112(t *testing.T) {

}
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil && root.Val == sum {
		return true
	} else if root.Left != nil && hasPathSum(root.Left, sum-(root.Val)) {
		return true
	} else if root.Right!=nil && hasPathSum(root.Right,sum-(root.Val)) {
		return true
	}
	return false
}

func Test118119(t *testing.T) {
	generate(6)
}

func generate(numRows int) [][]int {
	var a []int
	a = append(a, 1)
	var b [][]int
	if numRows==0 {
		return b
	}
	b= append(b, a)
	for i:=1;i< numRows;i++  {
		a=fuck118(a)
		b=append(b,a)
	}
	fmt.Println(b)
	return b
}
func getRow(rowIndex int) []int {
	var a []int
	a = append(a, 1)
	if rowIndex==0 {
		return a
	}
	for i:=1;i<= rowIndex;i++  {
		a=fuck118(a)
	}
	return a
}
func fuck118(hh []int)  []int{
	var gg []int
	for k,v := range hh {
		if k==0{
			gg = append(gg, v)
		}
		if k==(len(hh)-1){
			gg = append(gg, v)
			break
		}
		gg = append(gg, v+hh[k+1])
	}
	//	fmt.Println(gg)
	return gg
}