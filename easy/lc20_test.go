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