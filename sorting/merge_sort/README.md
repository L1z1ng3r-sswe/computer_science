# MergeSort

Here is the Go implementation of the **Merge Sort** algorithm along with a visualization of the sorting process.


## Merge Sort Algorithm

Merge sort is a classic divide-and-conquer algorithm that splits an array into halves, recursively sorts each half, and then merges the sorted halves back together.


## Usage

Here's an example demonstrating how to use the `merge_sort`:

```go
package main

import (
	"fmt"
	"github.com/L1z1ng3r-sswe/computer_science/sorting/merge_sort"
)

func main() {
	unSorted := []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println("sorted: ", merge_sort.MergeSort(unSorted)) // [1, 2, 3, 4, 5, 6, 7]
}
```

## Merge Sort Implementation

Here's the implementation of the merge sort algorithm:

```go
package merge_sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	sortedRightSide := MergeSort(nums[:mid])
	sortedLeftSide := MergeSort(nums[mid:])

	// order does not matter
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

	// order does not matter, because each of them is already sorted and one of them is empty
	sorted = append(sorted, nums1[i:]...)
	sorted = append(sorted, nums2[j:]...)
	return sorted
}
```

## Visualization

![Merge Sort Visualization](./assets/image.png)

The image above shows the binary tree visualization of the merge sort process. It illustrates how the array is split into smaller arrays and then merged back together in sorted order.