#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-04-21 23:00

"""Sol 1.8

Problem: Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

i.e. array contains positive + negative numbers

Approach: Kadane's Algo
Fact: 0 is always an answer in case of all numbers are negative (if subarray _could be empty_).
"""

from typing import List


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


class Solution:
    def max_subarray_sum(self, arr: List[int]) -> int:
        print(arr, end='\t')
        max_sum = float('-inf')  # or max_sum = 0
        curr_sum = 0

        for num in arr:
            curr_sum = max(num, curr_sum + num)
            max_sum = max(max_sum, curr_sum)

        print(max_sum)
        return max_sum


assert Solution().max_subarray_sum([-2, 1, -1, 3, -2, 4, -5]) == 5
assert Solution().max_subarray_sum([-2, -1, -1, -3, -2, -4, -5]) != 0
assert Solution().max_subarray_sum([-2, -1, -1, -3, -2, -4, -5]) == -1
