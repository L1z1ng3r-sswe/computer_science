package quick_sort

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	pivot := partition(nums)
	quickSort(nums[:pivot])
	quickSort(nums[pivot+1:])
}

func partition(nums []int) int {
	pvt := len(nums) - 1 // the last idx as a pivot and sort through it
	pvtVal := nums[pvt]
	left := 0

	for right := 0; right < len(nums)-1; right++ {
		if nums[right] < pvtVal {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	nums[left], nums[pvt] = nums[pvt], nums[left]
	return left
}
