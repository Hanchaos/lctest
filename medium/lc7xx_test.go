package medium

import (
	"fmt"
	"sort"
	"testing"
)

func Test739(t *testing.T) {
	a:=[]int{73, 74, 75, 71, 69, 72, 76, 73}
	b := dailyTemperatures(a)
	fmt.Println(b)
}

type Tem struct {
	Tem int
	index int
}
type TemStack struct {
	Data []Tem
}
func dailyTemperatures(T []int) []int {
	temStack := TemStack{Data:nil}
	dailyTem := make([]int, len(T))
	for i, val := range T {
		tem := Tem{
			Tem:   val,
			index: i,
		}
		if i == 0 {
			temStack.Push(tem)
			continue
		}
		for len(temStack.Data) > 0 {
			if temStack.Top().Tem < val {
				temp := temStack.Pop()
				dailyTem[temp.index] = i - temp.index
			}else{
				break
			}
		}
		temStack.Push(tem)
	}
	return dailyTem
}

func (this *TemStack) InitStack() TemStack {
	stack := TemStack{Data:nil}
	return stack
}

func (this *TemStack) Push( x Tem) {
	this.Data = append(this.Data, x)
}

func (this *TemStack) Pop() Tem{
	tem := this.Data[len(this.Data) - 1]
	this.Data = this.Data[:len(this.Data) - 1]
	return tem
}
func (this *TemStack) Top() Tem{
	return this.Data[len(this.Data) - 1]
}

func Test494(t *testing.T) {
	a:=[]int{1,1,1,1,1}
	fmt.Println(findTargetSumWays(a,3))
}
var count int
func findTargetSumWays(nums []int, S int) int {
	count=0
	help494(nums,0,S)
	return count
}
func help494(a []int, i , t int)  {
	if t==0 && i== len(a){
		count++
		return
	}
	if i==len(a) {return}
	t1:=t+a[i]
	help494(a,i+1,t1)
	t2:=t-a[i]
	help494(a,i+1,t2)
}

func Test438(t *testing.T) {
	findAnagrams("","bac")
}

func findAnagrams(s string, p string) []int {
	fmt.Println(p)

	return nil
}

func Test406(t *testing.T) {
	a:=[][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}
	fmt.Println(reconstructQueue(a))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)

	for i := 0; i < len(people); i++ {
		p := people[i][1]
		if p >= len(res) {
			res = append(res, people[i])
		} else {
			res = append(res, res[len(res)-1])
			copy(res[p+1:], res[p:])
			res[p] = people[i]
		}
	}

	return res
}

func rob(root *TreeNode) int {
	return Max(fuck337(root))
}
func fuck337(p1 *TreeNode)  int{
	if p1==nil{
		return 0
	}

	return Max(sum+p1.Val,sum+fuck337(p1.Left,sum)+fuck337(p1.Right,sum))
}