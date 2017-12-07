# -*- coding: utf-8 -*-
"""
Created on Thu Dec  7 16:27:44 2017

@author: toran.sahu
"""

def binary_search_recursive(arr,e):
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
    
    pivote = len(arr)//2 
    larr = arr[:pivote]
    rarr = arr[pivote:]
    if e >= arr[pivote]:
        return binary_search_recursive(rarr,e)
    else:
        return binary_search_recursive(larr,e)
    
arr = [1,2,3,4,5,6,7,8]    
print(binary_search_recursive(arr, 3))

def binary_search_iterative(arr,e):
    left = 0
    right = len(arr)-1
    found = False
    
    while(left<=right and not found):
        mid = (left + right)//2 
        if arr[mid] == e:
            found = True
        if e<arr[mid]:
            right = mid -1
        else:
            left = mid + 1
    return found

print(binary_search_iterative(arr, 2))