package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewHeap(t *testing.T) {
	maxHeap := NewHeap[int](func(i, j int) int {
		return j - i
	})
	minHeap := NewHeap[int](func(i, j int) int {
		return i - j
	})

	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 20; i++ {
		next := rand.Intn(1000)
		maxHeap.Push(next)
		minHeap.Push(next)
	}

	for maxHeap.Length() > 0 {
		maxHeap.Pop()
	}

	for _, i := range []int{328, 517, 660, 823, 981, 963, 356, 990, 582, 490, 919, 252, 282, 735, 873, 324, 640, 768, 39, 262, 577, 577} {
		maxHeap.Push(i)
	}
	maxHeap.Remove(577)
	maxHeap.Print()

	k := rand.Intn(1000)
	maxHeap.Push(k)
	minHeap.Push(k)
	maxHeap.Push(k)
	minHeap.Push(k)

	maxHeap.Print()
	minHeap.Print()

	fmt.Printf("remove k(%v)\n", k)
	maxHeap.Remove(k)
	minHeap.Remove(k)

	maxHeap.Print()
	minHeap.Print()

}
