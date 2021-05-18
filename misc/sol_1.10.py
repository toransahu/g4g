#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-04-21 23:00

"""Sol 1.10

Problem: Given an integer array nums, find the contiguous subarray which has the largest sum less than or equal to K.

i.e. array contains positive + negative numbers

Approach: Kadane's Algo
Fact: 0 is always an answer in case of all numbers are negative (if subarray _could be empty_).
"""

from typing import List


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


class Solution:
    # TODO: improve the time complexity
    def max_subarray_sum(self, arr: List[int], k: int) -> int:
        max_sum = float('-inf')  # or max_sum = 0
        for i, num in enumerate(arr):  # O(n)
            j = i
            curr_sum = 0
            while j < len(arr):  # O(n - 1)
                curr_sum += arr[j]
                if curr_sum == k:
                    max_sum = curr_sum
                    return max_sum
                if curr_sum > k:
                    break
                max_sum = max(max_sum, curr_sum)
                j += 1

        if max_sum < 0:
            max_sum = 0
        return max_sum


assert Solution().max_subarray_sum([-2, 1, -1, 3, -2, 4, -5], 5) == 5
assert Solution().max_subarray_sum([-2, 1, -1, 3, -2, 4, -5], 4) == 4
assert Solution().max_subarray_sum([-2, -1, -1, -3, -2, -4, -5], 4) == 0
