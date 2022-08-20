<h2 align="center">
    <img alt="LeetCode logo" src="../../docs/assets/leetcode.svg" />
    LeetCode • Easy problems • Implement strStr()
</h2>
<div align="center">

[![Profile](https://img.shields.io/badge/leetcode.com-smnvdev-f79f1b.svg?logo=leetcode)](https://leetcode.com/smnvdev/)
[![28. Implement strStr()](https://img.shields.io/badge/28._Implement_strStr()-00b8a3.svg?logo=leetcode)](https://leetcode.com/problems/implement-strstr/)
[![Submission Detail](https://img.shields.io/badge/Submission_Detail-Accepted-449d44.svg?logo=leetcode)](https://leetcode.com/submissions/detail/730414186/)
</div>

***

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1`
if `needle` is not part of `haystack`.

**Clarification:**

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's
[strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java's 
[indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).

**Example 1:**
```
Input: haystack = "hello", needle = "ll"
Output: 2
```
**Example 2:**
```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Constraints:**
 - `1 <= haystack.length, needle.length <= 10^4`
 - `haystack` and `needle` consist of only lowercase English characters.