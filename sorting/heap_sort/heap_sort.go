package heap_sort

func down(nums []int, n int, i int) { // for max heap
	for {
		ini := i
		rc := i*2 + 2
		lc := i*2 + 1

		if rc < n && nums[i] < nums[rc] {
			i = rc
		}

		if lc < n && nums[i] < nums[lc] {
			i = lc
		}

		if ini == i {
			return
		}
		nums[ini], nums[i] = nums[i], nums[ini]
	}
}

func heapSort(nums []int) {
	n := len(nums)

	// step 1: build a maxHeap
	for i := n/2 - 1; i >= 0; i-- { // half element to avoid redundant iterations
		down(nums, n, i)
	}

	// step 2: swap the largest with the smallest and heapify
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // move current largest to the end
		down(nums, i, 0)                    // pass i as n because [i:] now contains the largests nums
	}
}
