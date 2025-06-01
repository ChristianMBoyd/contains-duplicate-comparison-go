# contains-duplicate-comparison-go
## Introduction
Comparing the naive big-O analysis of LeetCode's "Contains Duplicate" with actual performant solutions.
## Problem statement
From https://leetcode.com/problems/contains-duplicate/description/ :
#### Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
The Go version of this LeetCode problem provides `nums` as an integer slice (`[]int`), which is subject to the following contraints:
#### 1 < `len(nums)` < 10<sup>5</sup>
#### -10<sup>9</sup> < `nums[i]` < 10<sup>9</sup> for any 0 <= `i` < `len(nums)`