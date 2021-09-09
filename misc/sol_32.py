#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
# created_on: 2021-08-05 15:57

"""Sol 32.

Way 1: find the pivot around which the array is rotated, then run binary_search on both the arrays.
How to find the pivot: in binary fashion (otherwise it would become O(n))

Way 2: combine finding pivot and then binary searching the array into single  step.  need to customize the standard binary search algorithm such that, until finding pivot, also check the key; if pivot is obtained, then run standard binary search on both the array.
"""


__author__ = 'Toran Sahu <toran.sahu@yahoo.com>'
__license__ = 'Distributed under terms of the MIT license'


def sol_32(arr, key):
    left_idx = 0
    right_idx = len(arr) - 1
    return binary_search(arr, left_idx, right_idx, key)


def binary_search(arr, left, right, key):
    if left < 0 or right >= len(arr) or left > right:
        return None

    # find the middle index
    mid_idx = (left + right) // 2
    # if key is the middle element
    if arr[mid_idx] == key:
        return mid_idx
    # else if arr from idx left to mid is sorted
    elif arr[left] <= arr[mid_idx]:
        # find the key there
        # if the key lies in left to mid range, then binary search
        if key >= arr[left] and key <= arr[mid_idx]:
            return binary_search(arr, left, mid_idx, key)
        # else if not found, then search the other part i.e. mid+1 to right
        else:
            return binary_search(arr, mid_idx+1, right, key)
    # else the mid+1 to right part MUST be sorted
    else:
        # find the key there
        # if the key lies in mid+1 to right range, then binary search
        if key >= arr[mid_idx+1] and key <= arr[right]:
            return binary_search(arr, mid_idx+1, right, key)
        # else if not found, then search the other part i.e. left to mid
        else:
            return binary_search(arr, left, mid_idx, key)
    return None


assert sol_32([5, 1, 2, 3, 4], 2) == 2
assert sol_32([2, 3, 4, 5, 1], 2) == 0
assert sol_32([2, 3, 4, 5, 1], 4) == 2
assert sol_32([2, 3, 4, 5, 1], 5) == 3
assert sol_32([2, 3, 4, 5, 1], 1) == 4
assert sol_32([1, 2, 3, 4, 5], 2) == 1
assert sol_32([5, 6, 7, 8, 9, 10, 1, 2, 3], 8) == 3
assert sol_32([5, 6, 7, 8, 9, 10, 1, 2, 3], 3) == 8
assert sol_32([5, 6, 7, 8, 9, 10, 1, 2, 3], 30) == None
assert sol_32([30, 40, 50, 10, 20], 10) == 3
