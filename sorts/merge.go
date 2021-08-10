package sorts

// Get len of list (n)
// Find middle point of list and split array
// if length of list != 1, repeat step 1 again for first half and second half
// Now we have list of 1 item each (sorted sublists), first element of both lists is compared, create new sorted list in ascending order,
// Then compared second item in both lists and add to new sorted list in ascending order
// Repeat step 2 until small sublists are empty

// 23 6 89 4 56 48

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
