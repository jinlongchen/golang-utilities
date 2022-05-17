package heap

import "fmt"

type Heap[T any] struct {
	arr        []T
	comparator func(i, j T) int
}

func NewHeap[T any](comparator func(i, j T) int) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
	}
}

func (h *Heap[T]) Peek() (T, bool) {
	if len(h.arr) == 0 {
		var result T
		return result, false
	}
	return h.arr[0], true
}

func (h *Heap[T]) MustPeek() T {
	val, ok := h.Peek()
	if !ok {
		panic("Overflow")
	}
	return val
}

func (h *Heap[T]) Length() int {
	return len(h.arr)
}

func (h *Heap[T]) Pop() (T, bool) {
	if len(h.arr) == 0 {
		var result T
		return result, false
	}

	v := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.down(0)
	return v, true
}

func (h *Heap[T]) MustPop() T {
	val, ok := h.Pop()
	if !ok {
		panic("Overflow")
	}
	return val
}

func (h *Heap[T]) Push(k T) bool {
	h.arr = append(h.arr, k)

	i := len(h.arr) - 1
	for i != 0 && h.comparator(h.arr[i], h.arr[(i-1)/2]) < 1 {
		h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
		i = (i - 1) / 2
	}
	return true
}

func (h *Heap[T]) Remove(k T) {
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

func (h *Heap[T]) Print() {
	for i := 0; i < 1; i++ {
		fmt.Printf("%v", h.arr[i])
	}
	for i := 1; i < h.Length(); i++ {
		fmt.Printf(", %v", h.arr[i])
	}
	fmt.Println()
}

func (h *Heap[T]) down(i int) {
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
