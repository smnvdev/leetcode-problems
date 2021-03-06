<h2 align="center">
    <img alt="LeetCode logo" src="../../docs/assets/leetcode.svg" />
    LeetCode • Easy problems • Merge Two Sorted Lists
</h2>
<div align="center">

[![Profile](https://img.shields.io/badge/leetcode.com-smnvdev-f79f1b.svg?logo=leetcode)](https://leetcode.com/smnvdev/)
[![21. Merge Two Sorted Lists](https://img.shields.io/badge/21._Merge_Two_Sorted_Lists-00b8a3.svg?logo=leetcode)](https://leetcode.com/problems/merge-two-sorted-lists/)
[![Submission Detail](https://img.shields.io/badge/Submission_Detail-Accepted-449d44.svg?logo=leetcode)](https://leetcode.com/submissions/detail/729574249/)
</div>

***

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return _the head of the merged linked list_.

**Example 1:**
```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**
```
Input: list1 = [], list2 = []
Output: []
```

**Example 3:**
```
Input: list1 = [], list2 = [0]
Output: [0]
```

**Constraints:**
 - The number of nodes in both lists is in the range `[0, 50]`.
 - `-100 <= Node.val <= 100`
 - Both `list1` and `list2` are sorted in non-decreasing order.