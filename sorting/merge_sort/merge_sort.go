package merge_sort

func mergeSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)

	for width := 1; width < n; width *= 2 {
		for i := 0; i < n; i += 2 * width {
			start := i
			mid := min(i+width, n)
			end := min(i+width*2, n)

			merge(nums, temp, start, mid, end)
		}
	}
}

func merge(nums []int, temp []int, start, mid, end int) { // end is exclusive
	k, i, j := start, start, mid

	for i < mid && j < end {
		if nums[i] > nums[j] {
			temp[k] = nums[j]
			j++
		} else {
			temp[k] = nums[i]
			i++
		}

		k++
	}

	for i < mid {
		temp[k] = nums[i]
		k++
		i++
	}

	for j < end {
		temp[k] = nums[j]
		k++
		j++
	}

	for i := start; i < end; i++ {
		nums[i] = temp[i]
	}
}
