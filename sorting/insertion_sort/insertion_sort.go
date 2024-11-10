package insertion_sort

func insertionSort(nums []int) {
	n := len(nums)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
