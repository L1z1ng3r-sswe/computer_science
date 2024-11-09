# InsertionSort

Here is the Go implementation of the **Insertion Sort** algorithm along with a detailed explanation of how the sorting process works.

## Insertion Sort Algorithm

Insertion sort is a simple sorting algorithm that works similarly to sorting playing cards in your hands. The array is virtually split into two parts: the sorted part and the unsorted part. The algorithm takes one element from the unsorted part, compares it with elements in the sorted part, and places it at the correct position.

## Insertion Sort Implementation


```go
func insertion_sort(nums []int) {
	n := len(nums)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
```

```cpp
void insertionSort(std::vector<int>& nums) {
  size_t n = nums.size();

  for (size_t i = 1; i < n; ++i) {
    for (size_t j = i; j > 0 && nums[j-1] > nums[j] ; --j ) {
      std::swap(nums[j], nums[j-1]);
    }
  }
}
```

### Time Complexity

Insertion sort has a time complexity of **O(n²)** in the worst and average cases, making it less efficient for large datasets. However, it performs well for small or nearly sorted datasets with a best-case time complexity of **O(n)**.

## Visualization

For a step-by-step visualization of the insertion sort algorithm, check out this video:

[![Insertion Sort Visualization](https://img.youtube.com/vi/IniSptbttgg/0.jpg)](https://www.youtube.com/watch?v=IniSptbttgg&ab_channel=GBhat)

The video linked above provides a visual explanation of how insertion sort works. You can see how elements are shifted and inserted at the correct positions.