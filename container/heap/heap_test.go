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
