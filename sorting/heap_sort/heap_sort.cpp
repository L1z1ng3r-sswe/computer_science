#include <iostream>
#include <vector>

void down(std::vector<int>& nums, size_t i, size_t n) { // for max-heap
  while (true) {
    size_t largest = i;
    size_t lc = i*2+1;
    size_t rc = i*2+2;

    if (lc < n && nums[largest] < nums[lc]) {
      largest = lc;
    }

    if (rc < n && nums[largest] < nums[rc]) {
      largest = rc;
    }

    if (largest == i) return ;
    std::swap(nums[i], nums[largest]);
    i = largest;
  }
}

void heapSort(std::vector<int>& nums) {
  size_t n = nums.size();

  // step 1: build max-heap
  for (size_t i = n/2-1; i >= 0; i--) { // start from the exact half to avoid redundant iterations
    down(nums, i, n);
  }

  // step 2: swap the largest with the smallest and heapify
  for (size_t i = n-1; i > 0; --i) { 
    std::swap(nums[0], nums[i]); // move current largest to the end
    down(nums, 0, i); // pass i as n because [i:] now contains the largests nums
  }
}