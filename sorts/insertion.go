package sorts

import "log"

// Get len of list (n)
// Start at pointer( pos = 0), this is start of sorted list
// Move to next item in unsorted list: position = j (start of unsorted list)
// Compare with predecessor and swap if not in correct order
// Move to j - 1, compare with predecessor and swap if not in correct order
// Repeat step 2steps above until reach start of list
// Increment pos +1
// Repeat inner loop steps above 3a-c
// Repeat steps 3 and 4 until pos == n

func InsertionSort(list []int32) {

	// Get len of list (n)
	length := len(list)

	// Start at pointer(position = 0), this is start of sorted list
	current := list[0]
	for i := 1; i < length; i++ {
		// Compare with predecessor and swap if not in correct order
		if list[i] < current {
			// Swap items
			tempPrev := list[i-1]
			list[i-1] = list[i]
			list[i] = tempPrev

			// Check already sorted items and put swapped item in the correct position
			for j := i; j >= 1; j-- {
				if list[j] < list[j-1] {
					prev := list[j-1]
					list[j-1] = list[j]
					list[j] = prev
				}
			}
		}
		current = list[i]
	}
	log.Println("sorted list via insertion sort: ", list)
}

// Alternate
// https://www.golangprograms.com/golang-program-for-implementation-of-insertionsort.html
func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
