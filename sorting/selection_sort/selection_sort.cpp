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