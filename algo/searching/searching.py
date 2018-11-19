# -*- coding: utf-8 -*-
"""
Created on Thu Dec  7 16:27:44 2017

@author: toran.sahu
"""

precision = 10


def ternary_search_recursive(arr, e, left, right):
    """
    Ternary search.
    Same as we do in binary search.
    Condition: Sorted array
    Equation:
    Time Complexity: O(log[3] n)
    Auxilary Space: O(log[3] n) recursive stack  space
    """

    if left < right:
        if right - left < precision:
            return linear_search(arr, e, left, right)
        print(left, right)
        pivot1 = (right - left) // 3
        pivot2 = (right - left) * 2 // 3

        if e == arr[pivot1]:
            return pivot1
        elif e == arr[pivot2]:
            return True
        elif e < arr[pivot1]:
            return ternary_search_recursive(arr, e, left, pivot1 - 1)
        elif e >= arr[pivot1] and e < arr[pivot2]:
            return ternary_search_recursive(arr, e, pivot1, pivot2 - 1)
        else:
            return ternary_search_recursive(arr, e, pivot2, right)
        return False
    return False


def exponetial_search(arr, e):
    """
    Enhancement over binary search.
    Enhancement: Find a range where element is present, and then do binary search found range.
    Range: 1..2..4..8..16...
    Condition: Sorted array
    Equation:
    Time Complexity: O(log n)
    Advantage: Better if array is infinite (Unbounded Searches/ Unbounded Binary Search)
    """

    size = 1
    while size * 2 <= len(arr):
        if arr[size - 1] >= e:
            break
        size *= 2
    # (i-1)/2 means we could not find a greater value in previous iteration
    range_start = int((size - 1) / 2)
    range_end = size - 1
    print(range_start, range_end)

    # binary search in range
    left = range_start
    right = range_end
    index = -1

    while left <= right and index == -1:
        mid = (left + right) // 2
        if arr[mid] == e:
            return mid
        if e < arr[mid]:
            right = mid - 1
        else:
            left = mid + 1
    return index


def interpolation_search(arr, e):
    """
    Enhancement over binary search.
    Enhancement: instead of take pivot as mid, it takes pivot by
    Pivot: left + ( (e-arr[left])*(right-left)/(arr[right]-arr[left]) )
    Condition: sorted array.
    Equation: 
    Time Complexity: O(log log n)
    """
    left = 0
    right = len(arr) - 1
    while left <= right and arr[left] <= e <= arr[right]:
        # pivot = left + int( (e-arr[left])*(right-left)/(arr[right]-arr[left]) )
        pivot = left + int(
            ((float(right - left) / (arr[right] - arr[left])) * (e - arr[left]))
        )
        if arr[pivot] == e:
            return pivot
        if e < arr[pivot]:
            right = pivot - 1
        else:
            left = pivot + 1
    return -1


def jump_search(arr, e):
    import math

    """
    Jump Search: Jumps a step and search
    Condition: sorted arrays.
    Similar to linear search.
    Complexity: O(sqrt(n))
    Best Value of m = sqrt(n)
    """
    m = int(math.sqrt(len(arr)))
    i = 0
    while i < len(arr) and e >= arr[i]:
        if e == arr[i]:
            return True
        i += m
    for j in range(i - m, min(len(arr), i)):
        if e == arr[j]:
            return True
    return False


def binary_search_iterative(arr, e):
    """
    Binary Search.
    Condition: Sorted array
    Complexity: log n
    """

    left = 0
    right = len(arr) - 1
    index = -1

    while left <= right and index == -1:
        mid = (left + right) // 2
        if arr[mid] == e:
            return mid
        if e < arr[mid]:
            right = mid - 1
        else:
            left = mid + 1
    return index


def binary_search_recursive(arr, e):
    """
    Binary Search.
    Condition: Sorted array
    Complexity: log n
    Equation: T(n) = T(n/2) + O(1)
    """

    if len(arr) == 1:
        if arr[0] == e:
            return True
        else:
            return False

    pivote = len(arr) // 2
    larr = arr[:pivote]
    rarr = arr[pivote:]
    if e >= arr[pivote]:
        return binary_search_recursive(rarr, e)
    else:
        return binary_search_recursive(larr, e)


def linear_search(arr, e, left, right):
    for i in range(left, right + 1):
        if arr[i] == e:
            return i
    return -1


arr = [1, 2, 3, 4, 5, 6, 7, 8]
# print(binary_search_recursive(arr, 3))
# print(binary_search_iterative(arr, 9))
# print(jump_search(arr, 1))
# print(interpolation_search(arr,9))
# print(exponetial_search(arr,8))
print(ternary_search_recursive(arr, 8, 0, len(arr) - 1))
