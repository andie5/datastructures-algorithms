package sorts

import (
	"fmt"
)

// Get len of list (n)
// Create a swap variable = false
// Start at beginning (pos=0)
// For current and next item, if not correct order swap them (set swap = true, only if swap done)
// Move to next adjacent items and check if they are in order, it not swap
// Repeat until you reach the end of list (n).
// If swap = true, set swap = false -> go to beginning of array and repeat process from step 3 else DONE

// 23 6 89 4 56 48
func BubbleSort(list []int32) {
	swap := true
	length := len(list)

	for swap {
		swap = false
		for i := 0; i < length-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swap = true
			}
		}
	}
	fmt.Println("sorted list via bubble sort: ", list)
}
