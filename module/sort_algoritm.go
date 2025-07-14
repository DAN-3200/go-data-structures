package module

import "fmt"

func RunWorkSpace() {
	array := []int{8, 7, 3, 10, 28, 4}
	BubbleSort(array)
}

func BubbleSort(box []int) {
	for y := 0; y < len(box)-1; y++ {
		for i := 0; i < len(box)-1-y; i++ {
			if box[i] > box[i+1] {
				box[i], box[i+1] = box[i+1], box[i]
			}
		}
	}

	fmt.Println(box)
}
