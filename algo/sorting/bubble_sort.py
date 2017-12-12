#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Dec 11 09:03:20 2017

@author: toran
"""

       
        
def insertion_sort(arr):
    """
    Select second & onwards elements one by one. ie. i
    Compare rest elements arr[j] i.e. 0...i-1 with above elements & if arr[j]
    is greater than arr[i] then shift jth elements right by one index.
    Repeat this for all j from i-1 to 0.
    
    Time Complexity: O(n2)
    Auxilary Space: 
    """
  
    n = len(arr)
    for i in range(1,n):
        k = i-1
        v = arr[i]
        for j in range(i-1,-1,-1):            
            if arr[j]>v:
                arr[j+1] = arr[j]
                k = j
        arr[k] = v        
    return arr
    

def selection_sort(arr):
    """
    Pulls a minimum element from sub-array and appends infront of that array.
    Time Complexity: 
        all case: O(n2)
    Auxilary Space: O(1)
    """
    n = len(arr)
    for i in range(0,n-1):
        min_idx = i
        for j in range(i,n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr

def bubble_sort_optimized(arr):
    """
    Optimized to detect an already sorted array.
    Time Complexity:
        Worst/Avg Case: O(n2)
        Best Case: O(n)
    """
    n = len(arr)
    swapped = False
    # travers the array n-1 times
    for i in range(0,n-1):
        # Last i elements are already in place
        for j in range(0,n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True
        if not swapped:
            return arr
    return arr


def bubble_sort(arr):
    """
    Time Complexity:
        Worst/Avg Case: O(n2)
        Best Case: O(n2)
    """
    n = len(arr)
    # travers the array n-1 times
    for i in range(0,n-1):
        # Last i elements are already in place
        for j in range(0,n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    return arr


arr = [3,6,9,1,4,9,0,3,5,2]
arr = [9,8,7,6,5,4,3,2,1]
arr_s = [1,2,3,4,5,5,6,7]
# print(bubble_sort(arr))
# print(bubble_sort_optimized(arr_s))
# print(selection_sort(arr))
print(insertion_sort(arr))
#print(insertionSort(arr))