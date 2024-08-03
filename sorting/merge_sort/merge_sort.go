package merge_sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	sortedRightSide := MergeSort(nums[:mid])
	sortedLeftSide := MergeSort(nums[mid:])

	// order is not matter
	return merge(sortedRightSide, sortedLeftSide)
}

func merge(nums1, nums2 []int) []int {
	var sorted []int
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}

	// order is not matter, because each of them are already sorted and one of them is empty
	sorted = append(sorted, nums1[i:]...)
	sorted = append(sorted, nums2[j:]...)
	return sorted
}
