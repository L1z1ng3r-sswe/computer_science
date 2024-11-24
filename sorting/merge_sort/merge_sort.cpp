#include <iostream>
#include <vector>

void merge(std::vector<int>& nums, std::vector<int>& temp, size_t start, size_t mid, size_t end) { // "end" is exclusive
  size_t i = start;
  size_t j = mid;
  size_t iterator = start;

  while (i < mid && j < end) {
    if (nums[i] > nums[j]) {
      temp[iterator++] = nums[j++];
    } else {
      temp[iterator++] = nums[i++];
    }
  }

  while (i < mid) {
    temp[iterator++] = nums[i++];
  }

  while (j < end) {
    temp[iterator++] = nums[j++];
  }

  for (size_t i = start; i < end; ++i) {
    nums[i] = temp[i];
  }
}

void mergeSort(std::vector<int>& nums) {
  size_t n = nums.size();
  std::vector<int> temp(n);

  for (size_t half = 1; half < n; half *= 2) {
    for (size_t i = 0; i < n; i += half * 2) {
      size_t start = i;
      size_t mid = std::min(i+half, n);
      size_t end = std::min(i+half*2, n);

      merge(nums, temp, start, mid, end);
    }
  }
}