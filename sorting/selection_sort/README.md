# SelectionSort

Here is the Go and C++ implementation of the **Selection sort** algorithm along with a step-by-step explanation of how the sorting process works.

## Selection Sort Algorithm

Selection sort is a simple sorting algorithm that divides the input array into two parts: the sorted part and the unsorted part. The algorithm repeatedly selects the smallest element from the unsorted part and swaps it with the first element of the unsorted part, thereby growing the sorted portion of the array.

## Selection Sort Implementation
```go
package main

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
```

```cpp
#include <iostream>
#include <vector>

void selectionSort(std::vector<int>& nums) {
  size_t n = nums.size();

  for (size_t i = 0; i < n-1; ++i) {
    int minIdx = i;

    for (size_t j = i+1; j < n; ++j) {
      if (nums[j] < nums[minIdx]) minIdx = j;
    }

    std::swap(nums[i], nums[minIdx]);
  }
}
```

## How Selection Sort Works

1. **Start at the beginning** of the array and consider the first element as the minimum.
2. **Search the unsorted portion** of the array for the smallest element.
3. **Swap the smallest element** found with the first element of the unsorted portion.
4. Move to the next element and repeat the process for the rest of the array.

This process continues until the entire array is sorted.

### Time Complexity

Selection sort has a time complexity of **O(n²)** due to the nested loops, where `n` is the number of elements in the array. As a result, this algorithm is inefficient for large datasets.

## Visualization

For a step-by-step visualization of the selection sort algorithm, check out this video:

[![Selection Sort Visualization](https://img.youtube.com/vi/MxEooU-8ps8/0.jpg)](https://www.youtube.com/watch?v=MxEooU-8ps8) 