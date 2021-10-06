package medium

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test208(t *testing.T) {
}
type Trie struct {
	isEnd bool
	next [26]*Trie
}
func Constructor() Trie {
	return Trie{}
}
func (this *Trie) Insert(word string) {
	node:=this
	for _, v := range word {
		index := v-'a'
		if node.next[index]==nil{
			trie:=Constructor()
			node.next[index] = &trie
		}
		node = node.next[index]
	}
	node.isEnd=true
}
func (this * Trie) Search(word string) bool {
	node := this
	for _, v := range word{
		index:=v-'a'
		node = node.next[index]
		if node==nil {return false}
	}
	return node.isEnd
}
func (this * Trie)  StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix{
		index:=v-'a'
		node = node.next[index]
		if node==nil {return false}
	}
	return true
}

func Test215(t *testing.T) {
	a:=[]int{3,2,1,4,5,8,6,7}
	fmt.Println(findKthLargest(a,3))
}

type S215 []int

func findKthLargest(nums []int, k int) int {
	h:=&S215{}
	heap.Init(h)
	for _,i := range nums {
		heap.Push(h,i)
		if len(*h) >k{
			heap.Pop(h)
		}
		fmt.Println(h)
	}
	fmt.Println(h)
	return heap.Pop(h).(int)
}

func (s S215) Len() int{
	return len(s)
}
func (s S215) Less(i,j int) bool {
	return s[i]>s[j]
}
func (s S215) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}
func (s *S215) Push(a interface{}) {
	*s = append(*s,a.(int))
}
func (s *S215) Pop() interface{}{
	old:=*s
	n:=len(old)
	x:=old[n-1]
	*s = old[0:n-1]
	return x
}

func Test1211(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b:=a[:5]
	b= append(b, 99)
	b[1]=20
	fmt.Println(a)
	fmt.Println(b)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==q || root==p || root==nil{
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	if left!=nil && right!=nil{
		return root
	}
	if left==nil {
		return right
	}else {
		return left}
}
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	i:= row-1
	j:=0
	for j<col && i>=0 {
		if matrix[i][j] > target{
			i--
		}else if matrix[i][j] <target{
			j++
		}else {
			return true
		}
	}
	return false
}

func moveZeroes(nums []int)  {
	l:=0
	r:=0
	for ;r<len(nums);r++  {
		if r!=0{
			nums[l]=nums[r]
			l++
		}
	}
	for ;l<len(nums);l++{
		nums[l]=0
	}
}

func Test000(t *testing.T) {
	n:=5
	var zbs []int
	for i:=0;i<n;i++{
		zb:=i
		fmt.Scanf("%d",&zb)
		zbs=append(zbs, zb)
	}
	fmt.Println(zbs)
}
func y000(a []int) []int{
	a[2]=100
	fmt.Printf("%p\n", a)
	a=append(a, 10)
	fmt.Printf("%p\n", a)
	return a
}