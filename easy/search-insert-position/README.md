<h2 align="center">
    <img alt="LeetCode logo" src="../../docs/assets/leetcode.svg" />
    LeetCode • Easy problems • Search Insert Position
</h2>
<div align="center">

[![Profile](https://img.shields.io/badge/leetcode.com-smnvdev-f79f1b.svg?logo=leetcode)](https://leetcode.com/smnvdev/)
[![35. Search Insert Position](https://img.shields.io/badge/35._Search_Insert_Position-00b8a3.svg?logo=leetcode)](https://leetcode.com/problems/search-insert-position/)
[![Submission Detail](https://img.shields.io/badge/Submission_Detail-Accepted-449d44.svg?logo=leetcode)](https://leetcode.com/submissions/detail/778448423/)
</div>

***

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return 
the index where it would be if it were inserted in order.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**
```
Input: nums = [1,3,5,6], target = 5
Output: 2
```

**Example 2:**
```
Input: nums = [1,3,5,6], target = 2
Output: 1
```

**Example 3:**
```
Input: nums = [1,3,5,6], target = 7
Output: 4
```


Constraints:

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` contains **distinct** values sorted in **ascending** order.
- `-10^4 <= target <= 10^4`