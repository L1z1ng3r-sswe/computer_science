  #include <iostream>
  #include <vector>
  #include <algorithm>


  size_t findPivot(std::vector<int>& nums, size_t start, size_t end) {
    size_t pvt = end;
    int pvtVal = nums[pvt];

    size_t left = start;

    for (size_t right = start; right < pvt; ++right) {
      if (nums[right] < pvtVal) {
        std::swap(nums[right], nums[left]); 
        left++;
      }
    } 

    std::swap(nums[left], nums[pvt]);

    return left;
  }

  void quickSort(std::vector<int>& nums, size_t start, size_t end) {
    if (end - start < 1) {
      return;
    }

    size_t pivot = findPivot(nums, start, end);

    if (pivot > 0) {
      quickSort(nums, start, pivot-1);
    } 
    quickSort(nums, pivot+1, end);
  }
