package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewHeapV1(t *testing.T) {
	maxHeap := NewHeapV1(func(i, j interface{}) int {
		return j.(int) - i.(int)
	})
	minHeap := NewHeap(func(i, j interface{}) int {
		return i.(int) - j.(int)
	})

	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 20; i++ {
		next := rand.Intn(1000)
		fmt.Printf("%v,", next)
		maxHeap.Push(next)
		minHeap.Push(next)
	}
	fmt.Println()

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
