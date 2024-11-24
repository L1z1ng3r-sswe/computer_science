# BubbleSort

Here is the Go and C++ implementation of the **Bubble Sort** algorithm along with a visualization of the sorting process.

## Bubble Sort Algorithm

Bubble sort is a simple comparison-based algorithm where adjacent elements are compared, and if they are in the wrong order, they are swapped. This process repeats until the entire array is sorted.


## Bubble Sort Implementation

```go
package bubble_sort

func bubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
```

```cpp
#include <iostream>
#include <vector>

void bubbleSort(std::vector<int>& nums) {
	size_t n = nums.size();
	for (size_t i = 0; i < n - 1; ++i) {
		for (size_t j = 0; j < n-i-1; ++j) {
			if (nums[j] > nums[j+1]){
				std::swap(nums[j], nums[j+1]);
			}
		}
	}
}
```

## Visualization

[![Bubble Sort Visualization](https://img.youtube.com/vi/0BkoXZBbhfU/0.jpg)](https://www.youtube.com/watch?v=0BkoXZBbhfU&ab_channel=GBhat)

## Files

- **C++ Implementation**: `bubble_sort.cpp`
- **Go Implementation**: `bubble_sort.go`