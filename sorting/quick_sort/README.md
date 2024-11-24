## Quick Sort Algorithm

Quick Sort works by partitioning the array around a "pivot" element and recursively sorting the partitions.

### Steps:
1. **Choose a Pivot**: Select an element from the array to act as the pivot. This could be the first, last, or a randomly chosen element.
2. **Partition the Array**: Rearrange the array such that:
   - Elements smaller than the pivot are placed on its left.
   - Elements larger than the pivot are placed on its right.
   - The pivot is then placed in its correct sorted position.
3. **Recursively Sort**:
   - Apply the same process to the left and right partitions.

---

## Key Features

1. **Divide and Conquer**: Splits the problem into smaller subproblems, solving each independently.
2. **In-Place**: The algorithm sorts the array without requiring significant additional memory.
3. **Efficient**: Performs well on average for large datasets.

---

## Advantages

1. **Fast on Average**: Quick Sort is generally faster than other \(O(n^2)\) algorithms like Selection Sort or Bubble Sort.
2. **Tailored for Large Datasets**: Performs well when dealing with large, random datasets.
3. **In-Place Sorting**: No need for extra memory except for the recursive stack.

---

## Limitations

1. **Worst-Case Time Complexity**: Degenerates to \(O(n^2)\) when the pivot leads to unbalanced partitions (e.g., already sorted data).
2. **Not Stable**: Does not preserve the order of equal elements.

---

## Complexity

### Time Complexity:
- **Best Case**: \(O(n \log n)\) — When the pivot divides the array into two nearly equal halves.
- **Average Case**: \(O(n \log n)\) — Expected performance with random input data.
- **Worst Case**: \(O(n^2)\) — When the pivot divides the array into highly unbalanced segments.

### Space Complexity:
- **In-Place Algorithm**: Requires \(O(\log n)\) additional space for recursion.

---

## Applications

1. **Sorting Large Arrays**: Ideal for large datasets due to its average-case efficiency.
2. **Systems with Limited Memory**: Performs well as an in-place algorithm.
3. **Divide-and-Conquer Paradigm**: Used in applications where dividing a problem into smaller subproblems is advantageous.

---

## Visualization

To see Quick Sort in action, watch this video for a step-by-step visual explanation:

[![Quick Sort Visualization](https://img.youtube.com/vi/Hoixgm4-P4M/0.jpg)](https://www.youtube.com/watch?v=Hoixgm4-P4M)

## Files

- **C++ Implementation**: `quick_sort.cpp`
- **Go Implementation**: `quick_sort.go`