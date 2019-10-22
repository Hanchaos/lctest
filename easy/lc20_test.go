package easy

import (
	"fmt"
	"testing"
)

func Test20(t *testing.T)  {
	a := "()[]{]}"
	isValid(a)
}

func fan(s string) string {
	switch s {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	case "}":
		return "{"
	case ")":
		return "("
	case "]":
		return "["
	}
	return ""
}


func isValid(s string) bool {
	ha := NewHaha()
	for _,v := range s{
		if ha.new() == fan(string(v)) {
			ha.pop()
		} else {
			ha.push(string(v))
		}
	}
	if len(ha.element)==0{
		fmt.Println("true")
		return true
	} else {
		fmt.Println("false")
		return false
	}
}

type haha struct {
	element []string
}

func NewHaha() *haha {
	return &haha{}
}

func (h *haha) push(s string) {
	h.element = append(h.element, s)
}
func (h *haha) pop() {
	size := len(h.element)
	if size == 0 {
		return
	}
	h.element = h.element[:size-1]
}
func (h *haha) new() string{
	size := len(h.element)
	if size == 0 {
		return ""
	}
	return h.element[size-1]
}

func Test21(t *testing.T){

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode
	result = new(ListNode)

	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}

	p1 := l1
	p2 := l2
	tmp := result

	for {
		if p1==nil || p2==nil{
			break
		}
		if p1.Val <= p2.Val {
			tmp.Next = p1
			tmp = tmp.Next
			p1 = p1.Next
		}else {
			tmp.Next = p2
			tmp = tmp.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		tmp.Next = p1
	}
	if p2 != nil {
		tmp.Next = p2
	}

	return result.Next // result self has 0
}

type ListNode struct {
	Val int
	Next *ListNode
}

func Test26(t *testing.T)  {
	a := []int{0,0,1,1,2,3,4,4,4}
	b := removeDuplicates(a)
	fmt.Println(a)
	println(b)
}

func removeDuplicates(nums []int) int {
	index := 0
	if len(nums)==0{
		return 0
	}
	tmp := nums[0]-1
	for _,v := range nums{
		if tmp != v {
			nums[index]=v
			tmp = v
			index++
		}
	}
	return index
}

func Test27(t *testing.T) {
	nums := []int{3,2,4,3,3,2,6,7,2,3}
	val := 3

	b := removeElement(nums,val)
	fmt.Println(nums)
	fmt.Println(b)
}

func removeElement(nums []int, val int) int {
	index := 0
	if len(nums)==0 {
		return 0
	}
	for _,v := range nums{
		if v != val {
			nums[index]=v
			index++
		}
	}
	return index
}


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