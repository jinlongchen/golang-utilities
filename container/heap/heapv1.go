package heap

import "fmt"

type HeapV1 struct {
	arr        []interface{}
	comparator func(i, j interface{}) int
}

func NewHeapV1(comparator func(i, j interface{}) int) *HeapV1 {
	return &HeapV1{
		comparator: comparator,
	}
}

func (h *HeapV1) Peek() (interface{}, bool) {
	if len(h.arr) == 0 {
		return nil, false
	}
	return h.arr[0], true
}

func (h *HeapV1) MustPeek() interface{} {
	val, ok := h.Peek()
	if !ok {
		panic("Overflow")
	}
	return val
}

func (h *HeapV1) Length() int {
	return len(h.arr)
}

func (h *HeapV1) Pop() (interface{}, bool) {
	if len(h.arr) == 0 {
		return nil, false
	}

	v := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.down(0)
	return v, true
}

func (h *HeapV1) MustPop() interface{} {
	val, ok := h.Pop()
	if !ok {
		panic("Overflow")
	}
	return val
}

func (h *HeapV1) Push(k interface{}) bool {
	h.arr = append(h.arr, k)

	i := len(h.arr) - 1
	for i != 0 && h.comparator(h.arr[i], h.arr[(i-1)/2]) < 1 {
		h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
		i = (i - 1) / 2
	}
	return true
}

func (h *HeapV1) Remove(k interface{}) {
	for h.Length() > 0 && h.comparator(h.arr[h.Length()-1], k) == 0 {
		h.arr = h.arr[:h.Length()-1]
	}
	for i := 0; i < h.Length(); i++ {
		if h.comparator(h.arr[i], k) == 0 {
			h.arr[i], h.arr[h.Length()-1] = h.arr[h.Length()-1], h.arr[i]
			h.arr = h.arr[:h.Length()-1]
			h.down(i)
		}
	}
}

func (h *HeapV1) Print() {
	for i := 0; i < 1; i++ {
		fmt.Printf("%v", h.arr[i])
	}
	for i := 1; i < h.Length(); i++ {
		fmt.Printf(", %v", h.arr[i])
	}
	fmt.Println()
}

func (h *HeapV1) down(i int) {
	left, right := (2*i)+1, (2*i)+2

	k := i

	if left < len(h.arr) && h.comparator(h.arr[left], h.arr[k]) < 1 {
		k = left
	}

	if right < len(h.arr) && h.comparator(h.arr[right], h.arr[k]) < 1 {
		k = right
	}

	if k != i {
		h.arr[i], h.arr[k] = h.arr[k], h.arr[i]
		h.down(k)
	}
}
