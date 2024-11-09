#include <iostream>
#include <vector>

void insertionSort(std::vector<int>& nums) {
  size_t n = nums.size();

  for (size_t i = 1; i < n; ++i) {
    for (size_t j = i; j > 0 && nums[j-1] > nums[j] ; --j ) {
      std::swap(nums[j], nums[j-1]);
    }
  }
}