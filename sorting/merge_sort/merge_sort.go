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

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	first := mergeSort(nums[:mid])
	second := mergeSort(nums[mid:])

	return merge(first, second)
}

// merge function receives 2 arrays and merge them into a single one
func merge(arr1, arr2 []int) []int {
	res := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)

	return res
}
