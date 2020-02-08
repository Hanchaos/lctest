package common


type HcStack struct {
	element []int
}
func NewStack() *HcStack {
	return &HcStack{}
}
func (h *HcStack) Push(s int) {
	h.element = append(h.element, s)
}
func (h *HcStack) Pop() {
	size := len(h.element)
	if size == 0 {
		return
	}
	h.element = h.element[:size-1]
}
func (h *HcStack) Top() int{
	size := len(h.element)
	if size == 0 {
		return 0
	}
	return h.element[size-1]
}