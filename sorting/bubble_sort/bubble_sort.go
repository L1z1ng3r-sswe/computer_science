package bubble_sort

func bubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := 1; j < n; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
