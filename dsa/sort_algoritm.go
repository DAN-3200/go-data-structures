package dsa

import (
	"fmt"
)

type Fibo struct {
	Count int
}

func (it *Fibo) Fibonacci(prev1, prev2 int) int {
	if it.Count == 19 {
		return prev1
	}
	
	newFibo := prev1 + prev2
	fmt.Println(newFibo)
	prev2, prev1 = prev1, newFibo
	it.Count++
	
	return it.Fibonacci(prev1, prev2)
}

func BubbleSort(box []int) []int {
	for y := 0; y < len(box)-1; y++ {
		for i := 0; i < len(box)-1-y; i++ {
			if box[i] > box[i+1] {
				box[i], box[i+1] = box[i+1], box[i]
			}
		}
	}

	return box
}
