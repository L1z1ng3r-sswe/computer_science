package selection_sort

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
